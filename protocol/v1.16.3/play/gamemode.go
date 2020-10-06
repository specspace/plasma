package play

import "github.com/spookspace/plasma/protocol/packet"

const (
	GamemodeSurvival  = packet.UnsignedByte(0x0)
	GamemodeCreative  = packet.UnsignedByte(0x1)
	GamemodeAdventure = packet.UnsignedByte(0x2)
	GamemodeSpectator = packet.UnsignedByte(0x3)
)
