package play

import (
	"github.com/specspace/plasma/protocol"
)

const (
	GamemodeSurvival  = protocol.UnsignedByte(0x0)
	GamemodeCreative  = protocol.UnsignedByte(0x1)
	GamemodeAdventure = protocol.UnsignedByte(0x2)
	GamemodeSpectator = protocol.UnsignedByte(0x3)
)
