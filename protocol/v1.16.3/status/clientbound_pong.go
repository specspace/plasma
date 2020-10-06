package status

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundPongPacketID byte = 0x01

type ClientBoundPong struct {
	Payload packet.Long
}

func (pk ClientBoundPong) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundPongPacketID,
		pk.Payload,
	)
}

func UnmarshalClientBoundPong(packet packet.Packet) (ClientBoundPong, error) {
	var pk ClientBoundPong

	if packet.ID != ClientBoundPongPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Payload,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
