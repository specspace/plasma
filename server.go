package plasma

import (
	"crypto/aes"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/specspace/plasma/protocol"
	"github.com/specspace/plasma/protocol/cfb8"
	"log"
	"net"
	"sync"
)

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeProtocol(w ResponseWriter, r *Request) {
	f(w, r)
}

type Handler interface {
	ServeProtocol(w ResponseWriter, r *Request)
}

const (
	DefaultServerAddr string = ":25565"
)

type muPlayers struct {
	sync.RWMutex
	players map[*conn]player
}

func (p *muPlayers) add(key *conn, value player) {
	p.Lock()
	defer p.Unlock()
	p.players[key] = value
}

func (p *muPlayers) get(key *conn) (player, error) {
	p.Lock()
	defer p.Unlock()
	pl, ok := p.players[key]
	if !ok {
		return player{}, errors.New("player does not exist")
	}
	return pl, nil
}

func (p *muPlayers) delete(key *conn) {
	p.Lock()
	defer p.Unlock()
	delete(p.players, key)
}

// Server defines the struct of a running Minecraft server
type Server struct {
	ID         string
	Addr       string
	Version    string
	ServerID   string
	Encryption bool

	SessionEncrypter     SessionEncrypter
	SessionAuthenticator SessionAuthenticator

	HandshakeHandler Handler
	StatusHandler    Handler
	LoginHandler     Handler
	PlayHandler      Handler

	listener  net.Listener
	players   muPlayers
	mu        sync.Mutex
	isRunning bool
}

func NewServerWithDefaults() (*Server, error) {
	encrypter, err := NewDefaultSessionEncrypter()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		Addr:                 DefaultServerAddr,
		Version:              "1.16.4",
		Encryption:           true,
		SessionEncrypter:     encrypter,
		SessionAuthenticator: &MojangSessionAuthenticator{},
		HandshakeHandler:     NewDefaultHandshakeMux(),
		StatusHandler:        NewDefaultStatusMux(),
		LoginHandler:         NewDefaultLoginMux(),
		PlayHandler:          disconnectHandler("You are now in Play state"),
		players: muPlayers{
			RWMutex: sync.RWMutex{},
			players: map[*conn]player{},
		},
		mu: sync.Mutex{},
	}, nil
}

func (srv *Server) AddPlayer(r *Request, username string) {
	srv.players.add(r.conn, player{
		conn:     r.conn,
		uuid:     uuid.NewV3(uuid.NamespaceOID, "OfflinePlayer:"+username),
		username: username,
	})
}

func (srv *Server) UpdatePlayer(r *Request, uuid uuid.UUID, skin Skin) error {
	player, err := srv.players.get(r.conn)
	if err != nil {
		return err
	}
	player.uuid = uuid
	player.skin = skin
	srv.players.add(r.conn, player)
	return nil
}

func (srv *Server) IsRunning() bool {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	return srv.isRunning
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
		addr = DefaultServerAddr
	}

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

		go srv.HandleConn(c)
	}
}

type ResponseWriter interface {
	PacketWriter
	NextState()
	EnableEncryption(sharedSecret []byte) error
	SetCompression(threshold int)
}

type responseWriter struct {
	nextState bool
	conn      *conn
}

func (w *responseWriter) NextState() {
	w.nextState = true
}

func (w *responseWriter) WritePacket(p protocol.Packet) error {
	return w.conn.WritePacket(p)
}

func (w *responseWriter) EnableEncryption(sharedSecret []byte) error {
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
	ProtocolState protocol.State
	Packet        protocol.Packet
	RemoteAddr    string
	Server        *Server
	Player        Player

	conn *conn
}

func (srv *Server) HandleConn(c net.Conn) {
	conn := wrapConn(c)
	defer conn.Close()
	srv.players.add(&conn, player{})
	defer srv.players.delete(&conn)

	var player Player

	for {
		pk, err := conn.ReadPacket()
		if err != nil {
			return
		}

		r := Request{
			ProtocolState: conn.State(),
			Packet:        pk,
			RemoteAddr:    conn.RemoteAddr().String(),
			Server:        srv,
			Player:        player,
			conn:          &conn,
		}

		w := responseWriter{
			conn: &conn,
		}

		switch conn.State() {
		case protocol.StateHandshaking:
			srv.HandshakeHandler.ServeProtocol(&w, &r)
			conn.state = protocol.StateStatus
			if w.nextState {
				conn.state = protocol.StateLogin
			}
		case protocol.StateStatus:
			srv.StatusHandler.ServeProtocol(&w, &r)
		case protocol.StateLogin:
			srv.LoginHandler.ServeProtocol(&w, &r)
			player, _ = srv.players.get(&conn)
			if w.nextState {
				conn.state = protocol.StatePlay
			}
		case protocol.StatePlay:
			srv.PlayHandler.ServeProtocol(&w, &r)
		}
	}
}
