package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundServerDifficultyPacketID byte = 0x0d

type ClientBoundServerDifficulty struct {
	Difficulty         protocol.UnsignedByte
	IsDifficultyLocked protocol.Boolean
}

func (pk ClientBoundServerDifficulty) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundServerDifficultyPacketID,
		pk.Difficulty,
		pk.IsDifficultyLocked,
	)
}

func UnmarshalClientBoundServerDifficulty(packet protocol.Packet) (ClientBoundServerDifficulty, error) {
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
