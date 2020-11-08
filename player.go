package plasma

import (
	"github.com/gofrs/uuid"
	"net"
)

type player struct {
	Conn
	uuid     uuid.UUID
	username string
	skin     Skin
}

type Player interface {
	// LocalAddr returns the local network address.
	LocalAddr() net.Addr

	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr

	// UUID returns the uuid of the player
	UUID() uuid.UUID

	// Username returns the username of the player
	Username() string

	// Username returns the Skin of the player
	Skin() Skin
}
