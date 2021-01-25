package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundPluginMessagePacketID byte = 0x17

type ClientBoundPluginMessage struct {
	Channel protocol.Identifier
	Data    protocol.OptionalByteArray
}

func (pk ClientBoundPluginMessage) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundPluginMessagePacketID,
		pk.Channel,
		pk.Data,
	)
}

func UnmarshalClientBoundPluginMessage(packet protocol.Packet) (ClientBoundPluginMessage, error) {
	var pk ClientBoundPluginMessage

	if packet.ID != ClientBoundPluginMessagePacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Channel,
		&pk.Data,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
