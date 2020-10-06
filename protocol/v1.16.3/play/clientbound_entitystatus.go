package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundEntityStatusPacketID byte = 0x17

type ClientBoundEntityStatus struct {
	EntityID     packet.Int
	EntityStatus packet.Byte
}

func (pk ClientBoundEntityStatus) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundEntityStatusPacketID,
		pk.EntityID,
		pk.EntityStatus,
	)
}

func UnmarshalClientBoundEntityStatus(packet packet.Packet) (ClientBoundEntityStatus, error) {
	var pk ClientBoundEntityStatus

	if packet.ID != ClientBoundEntityStatusPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.EntityID,
		&pk.EntityStatus,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
