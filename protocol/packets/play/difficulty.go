package play

import (
	"github.com/specspace/plasma/protocol"
)

const (
	DifficultyPeaceful = protocol.UnsignedByte(0x0)
	DifficultyEasy     = protocol.UnsignedByte(0x1)
	DifficultyNormal   = protocol.UnsignedByte(0x2)
	DifficultyHard     = protocol.UnsignedByte(0x3)
)
