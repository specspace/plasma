package plasma

import (
	"fmt"
	"github.com/specspace/plasma/protocol"
	"github.com/specspace/plasma/protocol/v1.16.4/handshaking"
	"github.com/specspace/plasma/protocol/v1.16.4/login"
	"github.com/specspace/plasma/protocol/v1.16.4/play"
	"github.com/specspace/plasma/protocol/v1.16.4/status"
	"sync"
)

type ServeMux struct {
	handlers map[byte]Handler
	mu       sync.RWMutex
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		handlers: map[byte]Handler{},
		mu:       sync.RWMutex{},
	}
}

func (mux *ServeMux) Handle(packetID byte, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if handler == nil {
		panic("plasma: nil handler")
	}

	mux.handlers[packetID] = handler
}

func (mux *ServeMux) HandleFunc(packetID byte, handler func(w ResponseWriter, r *Request)) {
	if handler == nil {
		panic("plasma: nil handler")
	}

	mux.Handle(packetID, HandlerFunc(handler))
}

func (mux *ServeMux) Handler(r *Request) (Handler, byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	handler, ok := mux.handlers[r.Packet.ID]
	if !ok {
		return HandlerFunc(panicHandler), r.Packet.ID
	}

	return handler, r.Packet.ID
}

func (mux *ServeMux) ServeProtocol(w ResponseWriter, r *Request) {
	handler, _ := mux.Handler(r)
	handler.ServeProtocol(w, r)
}

func panicHandler(w ResponseWriter, r *Request) {
	reason := fmt.Sprintf("%s(0x%x): Handler not implemented", r.ProtocolState, r.Packet.ID)
	disconnectHandler(reason)(w, r)
}

func disconnectHandler(reason string) HandlerFunc {
	return func(w ResponseWriter, r *Request) {
		if r.ProtocolState.IsLogin() {
			w.WritePacket(login.ClientBoundDisconnect{
				Reason: protocol.Chat(fmt.Sprintf("{\"text\":\"%s\"}", reason)),
			}.Marshal())
		} else if r.ProtocolState.IsPlay() {
			w.WritePacket(play.ClientBoundDisconnect{
				Reason: protocol.Chat(fmt.Sprintf("{\"text\":\"%s\"}", reason)),
			}.Marshal())
		}
	}
}

func NewDefaultHandshakeMux() *ServeMux {
	mux := NewServeMux()
	mux.HandleFunc(handshaking.ServerBoundHandshakePacketID, handshakeHandler)
	return mux
}

func handshakeHandler(w ResponseWriter, r *Request) {
	hs, err := handshaking.UnmarshalServerBoundHandshake(r.Packet)
	if err != nil {
		return
	}

	if hs.IsLoginRequest() {
		w.NextState()
	}
}

func NewDefaultStatusMux() *ServeMux {
	mux := NewServeMux()
	mux.HandleFunc(status.ServerBoundRequestPacketID, statusRequestHandler())
	mux.HandleFunc(status.ServerBoundPingPacketID, statusPingHandler)
	return mux
}

func statusRequestHandler() func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {
		w.WritePacket(status.ClientBoundResponse{
			JSONResponse: "{\"version\": {\"name\": \"1.16.4\",\"protocol\": 754},\"players\": {\"max\": 20,\"online\": 0,\"sample\": []},\"description\": {\"text\": \"Plasma lives\"},\"favicon\": \"data:image/png;base64,<data>\"}",
		}.Marshal())
	}
}

func statusPingHandler(w ResponseWriter, r *Request) {
	ping, err := status.UnmarshalServerBoundPing(r.Packet)
	if err != nil {
		return
	}

	w.WritePacket(status.ClientBoundPong{
		Payload: ping.Payload,
	}.Marshal())
}

func NewDefaultLoginMux() *ServeMux {
	mux := NewServeMux()
	mux.HandleFunc(login.ServerBoundLoginStartPacketID, loginStartHandler)
	mux.HandleFunc(login.ServerBoundEncryptionResponsePacketID, encryptionResponseHandler)
	return mux
}

func loginStartHandler(w ResponseWriter, r *Request) {
	s, err := login.UnmarshalServerBoundLoginStart(r.Packet)
	if err != nil {
		return
	}

	r.Server.AddPlayer(r, string(s.Name))

	if !r.Server.Encryption {
		disconnectHandler("Encryption required")(w, r)
		return
	}

	verifyToken, err := r.Server.SessionEncrypter.GenerateVerifyToken(r)
	if err != nil {
		disconnectHandler("Verify Token generation failed")(w, r)
		return
	}

	w.WritePacket(login.ClientBoundEncryptionRequest{
		ServerID:    protocol.String(r.Server.ID),
		PublicKey:   r.Server.SessionEncrypter.PublicKey(),
		VerifyToken: verifyToken,
	}.Marshal())
}

func encryptionResponseHandler(w ResponseWriter, r *Request) {
	resp, err := login.UnmarshalServerBoundEncryptionResponse(r.Packet)
	if err != nil {
		return
	}

	sharedSecret, err := r.Server.SessionEncrypter.DecryptAndVerifySharedSecret(r, resp.SharedSecret, resp.VerifyToken)
	if err != nil {
		disconnectHandler("Invalid encryption")(w, r)
		return
	}

	sessionHash := SessionHash(r.Server.ID, sharedSecret, r.Server.SessionEncrypter.PublicKey())
	session, err := r.Server.SessionAuthenticator.AuthenticateSession(r.Player.Username(), sessionHash)
	if err != nil {
		disconnectHandler("Invalid session")(w, r)
		return
	}

	if err := r.Server.UpdatePlayer(r, session.PlayerUUID, session.PlayerSkin); err != nil {
		return
	}

	if err := w.EnableEncryption(sharedSecret); err != nil {
		return
	}

	w.WritePacket(login.ClientBoundSetCompression{
		Threshold: 256,
	}.Marshal())

	w.SetCompression(256)

	w.WritePacket(login.ClientBoundLoginSuccess{
		UUID:     protocol.UUID(session.PlayerUUID),
		Username: protocol.String(r.Player.Username()),
	}.Marshal())
	w.WritePacket(play.ClientBoundDisconnect{Reason: "{\"text\":\"We made it to play mode!\"}"}.Marshal())
}
