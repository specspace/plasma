package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundSpawnExperienceOrbPacketID byte = 0x01

type ClientBoundSpawnExperienceOrb struct {
	EntityID protocol.VarInt
	X        protocol.Double
	Y        protocol.Double
	Z        protocol.Double
	Count    protocol.Short
}

func (pk ClientBoundSpawnExperienceOrb) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundSpawnExperienceOrbPacketID,
		pk.EntityID,
		pk.X,
		pk.Y,
		pk.Z,
		pk.Count,
	)
}

func UnmarshalClientBoundSpawnExperienceOrb(packet protocol.Packet) (ClientBoundSpawnExperienceOrb, error) {
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
