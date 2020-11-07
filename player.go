package plasma

import (
	"github.com/specspace/plasma/protocol"
)

type player struct {
	Conn
	username string
	skin     Skin
	state    protocol.State
}

type Player interface {
	Conn() Conn
	Username() string
	Skin() Skin
	State() protocol.State
}
