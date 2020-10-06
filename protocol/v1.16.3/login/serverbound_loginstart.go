package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ServerBoundLoginStartPacketID byte = 0x00

type ServerLoginStart struct {
	Name packet.String
}

func (pk ServerLoginStart) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundSetCompressionPacketID,
		pk.Name,
	)
}

func UnmarshalServerBoundLoginStart(packet packet.Packet) (ServerLoginStart, error) {
	var pk ServerLoginStart

	if packet.ID != ServerBoundLoginStartPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(&pk.Name); err != nil {
		return pk, err
	}

	return pk, nil
}
