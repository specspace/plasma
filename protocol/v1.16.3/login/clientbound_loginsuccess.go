package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundLoginSuccessPacketID byte = 0x02

type ClientBoundLoginSuccess struct {
	UUID     packet.UUID
	Username packet.String
}

func (pk ClientBoundLoginSuccess) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundLoginSuccessPacketID,
		pk.UUID,
		pk.Username,
	)
}

func ParseClientBoundLoginSuccess(packet packet.Packet) (ClientBoundLoginSuccess, error) {
	var pk ClientBoundLoginSuccess

	if packet.ID != ClientBoundLoginSuccessPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.UUID,
		&pk.Username,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
