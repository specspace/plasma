package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundSpawnExperienceOrbPacketID byte = 0x01

type ClientBoundSpawnExperienceOrb struct {
	EntityID packet.VarInt
	X        packet.Double
	Y        packet.Double
	Z        packet.Double
	Count    packet.Short
}

func (pk ClientBoundSpawnExperienceOrb) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundSpawnExperienceOrbPacketID,
		pk.EntityID,
		pk.X,
		pk.Y,
		pk.Z,
		pk.Count,
	)
}

func UnmarshalClientBoundSpawnExperienceOrb(packet packet.Packet) (ClientBoundSpawnExperienceOrb, error) {
	var pk ClientBoundSpawnExperienceOrb

	if packet.ID != ClientBoundSpawnExperienceOrbPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.EntityID,
		&pk.X,
		&pk.Y,
		&pk.Z,
		&pk.Count,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
