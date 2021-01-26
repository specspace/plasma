package plasma

import (
	"github.com/specspace/plasma/protocol"
	"sync"
)

type ServeMux struct {
	handlers map[protocol.State]map[byte]Handler
	mu       sync.RWMutex
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		handlers: map[protocol.State]map[byte]Handler{
			protocol.StateHandshaking: {},
			protocol.StateStatus:      {},
			protocol.StateLogin:       {},
			protocol.StatePlay:        {},
		},
		mu: sync.RWMutex{},
	}
}

func (mux *ServeMux) Handle(state protocol.State, packetID byte, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if handler == nil {
		panic("plasma: nil handler")
	}

	mux.handlers[state][packetID] = handler
}

func (mux *ServeMux) HandleFunc(state protocol.State, packetID byte, handler func(w ResponseWriter, r *Request)) {
	if handler == nil {
		panic("plasma: nil handler")
	}

	mux.Handle(state, packetID, HandlerFunc(handler))
}

func (mux *ServeMux) Handler(r *Request) (Handler, byte) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	handler, ok := mux.handlers[r.conn.state][r.Packet.ID]
	if !ok {
		return nil, r.Packet.ID
	}

	return handler, r.Packet.ID
}

func (mux *ServeMux) ServeProtocol(w ResponseWriter, r *Request) {
	handler, _ := mux.Handler(r)
	if handler == nil {
		return
	}

	handler.ServeProtocol(w, r)
}

var DefaultServeMux = NewServeMux()

func Handle(state protocol.State, packetID byte, handler Handler) {
	DefaultServeMux.Handle(state, packetID, handler)
}

func HandleFunc(state protocol.State, packetID byte, handler func(w ResponseWriter, r *Request)) {
	DefaultServeMux.HandleFunc(state, packetID, handler)
}
