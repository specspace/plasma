package plasma

import (
	"net"
)

const defaultAddr string = ":25565"

// Server defines the struct of a running Minecraft server
type Server struct {
	Addr        string
	listener    net.Listener
	isRunning   bool
	connections []Conn
}

func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = defaultAddr
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()
	srv.listener = l

	for {
		_, err := l.Accept()
		if err != nil {
			return err
		}
	}
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
