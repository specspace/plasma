package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundHeldItemChangePacketID byte = 0x40

type ClientBoundHeldItemChange struct {
	Slot packet.Byte
}

func (pk ClientBoundHeldItemChange) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundHeldItemChangePacketID,
		pk.Slot,
	)
}

func UnmarshalClientBoundHeldItemChange(packet packet.Packet) (ClientBoundHeldItemChange, error) {
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
