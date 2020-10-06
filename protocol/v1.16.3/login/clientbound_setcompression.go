package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundSetCompressionPacketID byte = 0x03

type ClientBoundSetCompression struct {
	Threshold packet.VarInt
}

func (pk ClientBoundSetCompression) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundSetCompressionPacketID,
		pk.Threshold,
	)
}

func ParseClientBoundSetCompression(packet packet.Packet) (ClientBoundSetCompression, error) {
	var pk ClientBoundSetCompression

	if packet.ID != ClientBoundSetCompressionPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(&pk.Threshold); err != nil {
		return pk, err
	}

	return pk, nil
}
