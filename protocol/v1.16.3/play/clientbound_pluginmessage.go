package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundPluginMessagePacketID byte = 0x17

type ClientBoundPluginMessage struct {
	Channel packet.Identifier
	Data    packet.OptionalByteArray
}

func (pk ClientBoundPluginMessage) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundPluginMessagePacketID,
		pk.Channel,
		pk.Data,
	)
}

func UnmarshalClientBoundPluginMessage(packet packet.Packet) (ClientBoundPluginMessage, error) {
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
