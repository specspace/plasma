package plasma

import (
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
	if handler == nil {
		return
	}

	handler.ServeProtocol(w, r)
}
