package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundPlayerAbilitiesPacketID byte = 0x30

type ClientBoundPlayerAbilities struct {
	Flags               protocol.Byte
	FlyingSpeed         protocol.Float
	FieldOfViewModifier protocol.Float
}

func (pk ClientBoundPlayerAbilities) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundPlayerAbilitiesPacketID,
		pk.Flags,
		pk.FlyingSpeed,
		pk.FieldOfViewModifier,
	)
}

func UnmarshalClientBoundPlayerAbilities(packet protocol.Packet) (ClientBoundPlayerAbilities, error) {
	var pk ClientBoundPlayerAbilities

	if packet.ID != ClientBoundPlayerAbilitiesPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Flags,
		&pk.FlyingSpeed,
		&pk.FieldOfViewModifier,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
