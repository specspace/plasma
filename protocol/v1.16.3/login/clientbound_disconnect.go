package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundDisconnectPacketID byte = 0x00

type ClientBoundDisconnect struct {
	Reason packet.Chat
}

func (pk ClientBoundDisconnect) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundDisconnectPacketID,
		pk.Reason,
	)
}

func UnmarshalClientBoundDisconnect(packet packet.Packet) (ClientBoundDisconnect, error) {
	var pk ClientBoundDisconnect

	if packet.ID != ClientBoundDisconnectPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Reason,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
