package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundLoginPluginRequestPacketID byte = 0x04

type ClientBoundLoginPluginRequest struct {
	MessageID packet.VarInt
	Channel   packet.Identifier
	Data      packet.OptionalByteArray
}

func (pk ClientBoundLoginPluginRequest) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundLoginPluginRequestPacketID,
		pk.MessageID,
		pk.Channel,
		pk.Data,
	)
}

func UnmarshalClientBoundLoginPluginRequest(packet packet.Packet) (ClientBoundLoginPluginRequest, error) {
	var pk ClientBoundLoginPluginRequest

	if packet.ID != ClientBoundLoginPluginRequestPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.MessageID,
		&pk.Channel,
		&pk.Data,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
