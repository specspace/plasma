package plasma

import (
	"github.com/gofrs/uuid"
)

type player struct {
	*conn
	uuid     uuid.UUID
	username string
	skin     Skin
}

func (p player) UUID() uuid.UUID {
	return p.uuid
}

func (p player) Username() string {
	return p.username
}

func (p player) Skin() Skin {
	return p.skin
}

type Player interface {
	Conn

	// UUID returns the uuid of the player
	UUID() uuid.UUID

	// Username returns the username of the player
	Username() string

	// Skin returns the Skin of the player
	Skin() Skin
}
