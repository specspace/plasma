package play

import "github.com/spookspace/plasma/protocol/packet"

const (
	DifficultyPeaceful = packet.UnsignedByte(0x0)
	DifficultyEasy     = packet.UnsignedByte(0x1)
	DifficultyNormal   = packet.UnsignedByte(0x2)
	DifficultyHard     = packet.UnsignedByte(0x3)
)
