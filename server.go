package plasma

import (
	"github.com/specspace/plasma/protocol"
	"net"
)

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeProtocol(w ResponseWriter, r *Request) {
	f(w, r)
}

type Handler interface {
	ServeProtocol(w ResponseWriter, r *Request)
}

const DefaultServerAddr string = ":25565"

// Server defines the struct of a running Minecraft server
type Server struct {
	Addr                 string
	Version              string
	Encryption           bool
	SessionAuthenticator SessionAuthenticator
	HandshakeHandler     Handler
	StatusHandler        Handler
	LoginHandler         Handler
	PlayHandler          Handler

	listener    net.Listener
	isRunning   bool
	connections []Conn
}

func (srv Server) IsRunning() bool {
	return srv.isRunning
}

func (srv *Server) Close() error {
	if srv.listener != nil {
		return srv.listener.Close()
	}

	return nil
}

func (srv *Server) ListenAndServe() error {
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
}

type responseWriter struct {
	packet    *protocol.Packet
	nextState bool
}

func (w *responseWriter) WritePacket(p protocol.Packet) error {
	w.packet = &p
	return nil
}

func (w *responseWriter) NextState() {
	w.nextState = true
}

type Request struct {
	ProtocolState protocol.State
	Packet        protocol.Packet
	RemoteAddr    string
}

func (srv Server) HandleConn(c net.Conn) {
	conn := wrapConn(c)
	defer conn.Close()

	for {
		pk, err := conn.ReadPacket()
		if err != nil {
			return
		}

		r := Request{
			ProtocolState: conn.State(),
			Packet:        pk,
			RemoteAddr:    conn.RemoteAddr().String(),
		}

		w := responseWriter{}

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
			if w.nextState {
				conn.state = protocol.StatePlay
			}
		case protocol.StatePlay:
			srv.PlayHandler.ServeProtocol(&w, &r)
		}

		if w.packet == nil {
			continue
		}

		if err := conn.WritePacket(*w.packet); err != nil {
			return
		}
	}
}
