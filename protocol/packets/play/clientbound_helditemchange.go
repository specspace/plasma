package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundHeldItemChangePacketID byte = 0x40

type ClientBoundHeldItemChange struct {
	Slot protocol.Byte
}

func (pk ClientBoundHeldItemChange) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundHeldItemChangePacketID,
		pk.Slot,
	)
}

func UnmarshalClientBoundHeldItemChange(packet protocol.Packet) (ClientBoundHeldItemChange, error) {
	var pk ClientBoundHeldItemChange

	if packet.ID != ClientBoundHeldItemChangePacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Slot,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
