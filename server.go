package plasma

import (
	"bufio"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/specspace/plasma/protocol"
	"github.com/specspace/plasma/protocol/cfb8"
	"net"
	"os"
	"sync"
)

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeProtocol(w ResponseWriter, r *Request) {
	f(w, r)
}

type Handler interface {
	ServeProtocol(w ResponseWriter, r *Request)
}

type StatusResponse struct {
	DisconnectMessage string
	Version           Version
	IconPath          string
	Motd              string
	MaxPlayers        int
	PlayersOnline     int
	Players           []struct {
		Name string
		ID   string
	}
}

type StatusResponseJSON struct {
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"sample"`
	} `json:"players"`
	Description struct {
		Text string `json:"text"`
	} `json:"description"`
	Favicon string `json:"favicon"`
}

func (sr StatusResponse) JSON() (StatusResponseJSON, error) {
	var response StatusResponseJSON
	response.Version.Name = sr.Version.Name
	response.Version.Protocol = sr.Version.ProtocolNumber
	response.Players.Max = sr.MaxPlayers
	response.Players.Online = sr.PlayersOnline
	response.Description.Text = sr.Motd

	for _, p := range sr.Players {
		response.Players.Sample = append(response.Players.Sample,
			struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			}{
				Name: p.Name,
				ID:   p.ID,
			},
		)
	}

	if sr.IconPath != "" {
		img64, err := loadImageAndEncodeToBase64String(sr.IconPath)
		if err != nil {
			return response, err
		}
		response.Favicon = fmt.Sprintf("data:image/png;base64,%s", img64)
	}

	return response, nil
}

func loadImageAndEncodeToBase64String(path string) (string, error) {
	if path == "" {
		return "", nil
	}

	imgFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	fileInfo, err := imgFile.Stat()
	if err != nil {
		return "", err
	}

	buffer := make([]byte, fileInfo.Size())
	fileReader := bufio.NewReader(imgFile)
	_, err = fileReader.Read(buffer)
	if err != nil {
		return "", nil
	}

	return base64.StdEncoding.EncodeToString(buffer), nil
}

const DefaultAddr string = ":25565"

// Server defines the struct of a running Minecraft server
type Server struct {
	ID         string
	Addr       string
	Encryption bool

	SessionEncrypter     SessionEncrypter
	SessionAuthenticator SessionAuthenticator
	Handler              Handler

	listener net.Listener
	players  map[*conn]player
	mu       sync.RWMutex
}

func (srv *Server) getPlayer(conn *conn) player {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	return srv.players[conn]
}

func (srv *Server) putPlayer(c *conn, p player) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	srv.players[c] = p
}

func (srv *Server) removePlayer(c *conn) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	delete(srv.players, c)
}

func (srv *Server) Player(r *Request) Player {
	return srv.getPlayer(r.conn)
}

func (srv *Server) Players() []Player {
	srv.mu.RLock()
	defer srv.mu.RUnlock()

	var players []Player
	for _, player := range srv.players {
		players = append(players, player)
	}
	return players
}

func (srv *Server) AddPlayer(r *Request, username string) {
	srv.putPlayer(r.conn, player{
		conn:     r.conn,
		uuid:     uuid.NewV3(uuid.NamespaceOID, "OfflinePlayer:"+username),
		username: username,
	})
}

func (srv *Server) IsRunning() bool {
	return srv.listener != nil
}

func (srv *Server) Close() error {
	if srv.listener != nil {
		return srv.listener.Close()
	}

	return nil
}

func (srv *Server) ListenAndServe() error {
	if srv.listener != nil {
		return errors.New("server is already running")
	}

	addr := srv.Addr
	if addr == "" {
		addr = DefaultAddr
	}

	srv.players = map[*conn]player{}
	srv.mu = sync.RWMutex{}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()
	srv.listener = l

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		go srv.serve(c)
	}
}

func (srv *Server) serve(c net.Conn) {
	conn := wrapConn(c)
	srv.putPlayer(conn, player{})
	defer func() {
		srv.removePlayer(conn)
		conn.Close()
	}()

	r := Request{
		RemoteAddr: conn.RemoteAddr().String(),
		Server:     srv,
		conn:       conn,
	}

	w := responseWriter{
		conn: conn,
	}

	for {
		pk, err := conn.ReadPacket()
		if err != nil {
			return
		}

		r.Packet = pk

		srv.Handler.ServeProtocol(&w, &r)
	}
}

func ListenAndServe(addr string, handler Handler) error {
	encrypter, err := NewDefaultSessionEncrypter()
	if err != nil {
		return err
	}

	srv := &Server{
		Addr:                 addr,
		Encryption:           true,
		SessionEncrypter:     encrypter,
		SessionAuthenticator: &MojangSessionAuthenticator{},
		Handler:              handler,
	}

	return srv.ListenAndServe()
}

type ResponseWriter interface {
	PacketWriter
	SetState(state protocol.State)
	SetEncryption(sharedSecret []byte) error
	SetCompression(threshold int)
}

type responseWriter struct {
	nextState bool
	conn      *conn
}

func (w *responseWriter) SetState(state protocol.State) {
	w.conn.state = state
}

func (w *responseWriter) WritePacket(p protocol.Packet) error {
	return w.conn.WritePacket(p)
}

func (w *responseWriter) SetEncryption(sharedSecret []byte) error {
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return err
	}

	w.conn.SetCipher(
		cfb8.NewEncrypter(block, sharedSecret),
		cfb8.NewDecrypter(block, sharedSecret),
	)
	return nil
}

func (w *responseWriter) SetCompression(threshold int) {
	w.conn.threshold = threshold
}

type Request struct {
	Packet     protocol.Packet
	RemoteAddr string
	Server     *Server

	conn *conn
}

func (r Request) ProtocolState() protocol.State {
	return r.conn.state
}

func (r *Request) Player() Player {
	return r.Server.Player(r)
}

func (r *Request) UpdatePlayerUsername(username string) {
	player := r.Server.getPlayer(r.conn)
	player.username = username
	r.Server.putPlayer(r.conn, player)
}

func (r *Request) UpdatePlayerUUID(uuid uuid.UUID) {
	player := r.Server.getPlayer(r.conn)
	player.uuid = uuid
	r.Server.putPlayer(r.conn, player)
}

func (r *Request) UpdatePlayerSkin(skin Skin) {
	player := r.Server.getPlayer(r.conn)
	player.skin = skin
	r.Server.putPlayer(r.conn, player)
}
