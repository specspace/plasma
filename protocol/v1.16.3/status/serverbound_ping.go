package status

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundPingPacketID byte = 0x01

type ClientBoundPing struct {
	Payload packet.Long
}

func (pk ClientBoundPing) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundPingPacketID,
		pk.Payload,
	)
}

func UnmarshalClientBoundPing(packet packet.Packet) (ClientBoundPing, error) {
	var pk ClientBoundPing

	if packet.ID != ClientBoundPingPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Payload,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
