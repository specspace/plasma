package plasma

import (
	"github.com/specspace/plasma/protocol/v1.16.4/handshaking"
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
		return nil, r.Packet.ID
	}

	return handler, r.Packet.ID
}

func (mux *ServeMux) ServeProtocol(w ResponseWriter, r *Request) {
	handler, _ := mux.Handler(r)
	handler.ServeProtocol(w, r)
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
	mux.HandleFunc(status.ServerBoundRequestPacketID, requestHandler())
	mux.HandleFunc(status.ServerBoundPingPacketID, pingHandler)
	return mux
}

func requestHandler() func(w ResponseWriter, r *Request) {
	return func(w ResponseWriter, r *Request) {
		w.WritePacket(status.ClientBoundResponse{
			JSONResponse: "{\"version\": {\"name\": \"1.16.4\",\"protocol\": 754},\"players\": {\"max\": 20,\"online\": 0,\"sample\": []},\"description\": {\"text\": \"Plasma lives\"},\"favicon\": \"data:image/png;base64,<data>\"}",
		}.Marshal())
	}
}

func pingHandler(w ResponseWriter, r *Request) {
	ping, err := status.UnmarshalServerBoundPing(r.Packet)
	if err != nil {
		return
	}

	w.WritePacket(status.ClientBoundPong{
		Payload: ping.Payload,
	}.Marshal())
}
