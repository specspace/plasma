package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundPlayerAbilitiesPacketID byte = 0x30

type ClientBoundPlayerAbilities struct {
	Flags               packet.Byte
	FlyingSpeed         packet.Float
	FieldOfViewModifier packet.Float
}

func (pk ClientBoundPlayerAbilities) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundPlayerAbilitiesPacketID,
		pk.Flags,
		pk.FlyingSpeed,
		pk.FieldOfViewModifier,
	)
}

func UnmarshalClientBoundPlayerAbilities(packet packet.Packet) (ClientBoundPlayerAbilities, error) {
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
