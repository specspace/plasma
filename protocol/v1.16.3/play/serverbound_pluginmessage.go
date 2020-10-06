package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ServerBoundPluginMessagePacketID byte = 0x17

type ServerBoundPluginMessage struct {
	Channel packet.Identifier
	Data    packet.OptionalByteArray
}

func (pk ServerBoundPluginMessage) Marshal() packet.Packet {
	return packet.Marshal(
		ServerBoundPluginMessagePacketID,
		pk.Channel,
		pk.Data,
	)
}

func UnmarshalServerBoundPluginMessage(packet packet.Packet) (ServerBoundPluginMessage, error) {
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
