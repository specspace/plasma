package play

import (
	"github.com/specspace/plasma/protocol"
)

const ServerBoundClientSettingsPacketID byte = 0x05

type ServerBoundClientSettings struct {
	Locale             protocol.String
	ViewDistance       protocol.Byte
	ChatMode           protocol.VarInt
	ChatColors         protocol.Boolean
	DisplayedSkinParts protocol.UnsignedByte
	MainHand           protocol.VarInt
}

func (pk ServerBoundClientSettings) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ServerBoundClientSettingsPacketID,
		pk.Locale,
		pk.ViewDistance,
		pk.ChatMode,
		pk.ChatColors,
		pk.DisplayedSkinParts,
		pk.MainHand,
	)
}

func UnmarshalServerBoundClientSettings(packet protocol.Packet) (ServerBoundClientSettings, error) {
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
