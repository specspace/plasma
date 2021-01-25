package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundEntityStatusPacketID byte = 0x17

type ClientBoundEntityStatus struct {
	EntityID     protocol.Int
	EntityStatus protocol.Byte
}

func (pk ClientBoundEntityStatus) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundEntityStatusPacketID,
		pk.EntityID,
		pk.EntityStatus,
	)
}

func UnmarshalClientBoundEntityStatus(packet protocol.Packet) (ClientBoundEntityStatus, error) {
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
