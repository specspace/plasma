package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundServerDifficultyPacketID byte = 0x0d

type ClientBoundServerDifficulty struct {
	Difficulty         packet.UnsignedByte
	IsDifficultyLocked packet.Boolean
}

func (pk ClientBoundServerDifficulty) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundServerDifficultyPacketID,
		pk.Difficulty,
		pk.IsDifficultyLocked,
	)
}

func UnmarshalClientBoundServerDifficulty(packet packet.Packet) (ClientBoundServerDifficulty, error) {
	var pk ClientBoundServerDifficulty

	if packet.ID != ClientBoundServerDifficultyPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Difficulty,
		&pk.IsDifficultyLocked,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
