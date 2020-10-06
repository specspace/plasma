package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ServerBoundLoginPluginResponsePacketID byte = 0x04

type ServerBoundLoginPluginResponse struct {
	MessageID  packet.VarInt
	Successful packet.Boolean
	Data       packet.OptionalByteArray
}

func (pk ServerBoundLoginPluginResponse) Marshal() packet.Packet {
	return packet.Marshal(
		ServerBoundLoginPluginResponsePacketID,
		pk.MessageID,
		pk.Successful,
		pk.Data,
	)
}

func UnmarshalServerBoundLoginPluginResponse(packet packet.Packet) (ServerBoundLoginPluginResponse, error) {
	var pk ServerBoundLoginPluginResponse

	if packet.ID != ServerBoundLoginPluginResponsePacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.MessageID,
		&pk.Successful,
		&pk.Data,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
