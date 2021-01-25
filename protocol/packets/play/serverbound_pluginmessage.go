package play

import (
	"github.com/specspace/plasma/protocol"
)

const ServerBoundPluginMessagePacketID byte = 0x17

type ServerBoundPluginMessage struct {
	Channel protocol.Identifier
	Data    protocol.OptionalByteArray
}

func (pk ServerBoundPluginMessage) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ServerBoundPluginMessagePacketID,
		pk.Channel,
		pk.Data,
	)
}

func UnmarshalServerBoundPluginMessage(packet protocol.Packet) (ServerBoundPluginMessage, error) {
	var pk ServerBoundPluginMessage

	if packet.ID != ServerBoundPluginMessagePacketID {
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
