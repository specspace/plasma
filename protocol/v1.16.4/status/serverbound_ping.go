package status

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundPingPacketID byte = 0x01

type ClientBoundPing struct {
	Payload protocol.Long
}

func (pk ClientBoundPing) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundPingPacketID,
		pk.Payload,
	)
}

func UnmarshalClientBoundPing(packet protocol.Packet) (ClientBoundPing, error) {
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
