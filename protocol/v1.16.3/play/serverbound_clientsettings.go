package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ServerBoundClientSettingsPacketID byte = 0x05

type ServerBoundClientSettings struct {
	Locale             packet.String
	ViewDistance       packet.Byte
	ChatMode           packet.VarInt
	ChatColors         packet.Boolean
	DisplayedSkinParts packet.UnsignedByte
	MainHand           packet.VarInt
}

func (pk ServerBoundClientSettings) Marshal() packet.Packet {
	return packet.Marshal(
		ServerBoundClientSettingsPacketID,
		pk.Locale,
		pk.ViewDistance,
		pk.ChatMode,
		pk.ChatColors,
		pk.DisplayedSkinParts,
		pk.MainHand,
	)
}

func UnmarshalServerBoundClientSettings(packet packet.Packet) (ServerBoundClientSettings, error) {
	var pk ServerBoundClientSettings

	if packet.ID != ServerBoundClientSettingsPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Locale,
		&pk.ViewDistance,
		&pk.ChatMode,
		&pk.ChatColors,
		&pk.DisplayedSkinParts,
		&pk.MainHand,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
