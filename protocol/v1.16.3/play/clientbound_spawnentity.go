package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundSpawnEntityPacketID byte = 0x00

type ClientBoundSpawnEntity struct {
	EntityID   packet.VarInt
	ObjectUUID packet.UUID
	Type       packet.VarInt
	X          packet.Double
	Y          packet.Double
	Z          packet.Double
	Pitch      packet.Angle
	Yaw        packet.Angle
	VelocityX  packet.Short
	VelocityY  packet.Short
	VelocityZ  packet.Short
}

func (pk ClientBoundSpawnEntity) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundSpawnEntityPacketID,
		pk.EntityID,
		pk.ObjectUUID,
		pk.Type,
		pk.X,
		pk.Y,
		pk.Z,
		pk.Pitch,
		pk.Yaw,
		pk.VelocityX,
		pk.VelocityY,
		pk.VelocityZ,
	)
}

func UnmarshalClientBoundSpawnEntity(packet packet.Packet) (ClientBoundSpawnEntity, error) {
	var pk ClientBoundSpawnEntity

	if packet.ID != ClientBoundSpawnEntityPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.EntityID,
		&pk.ObjectUUID,
		&pk.Type,
		&pk.X,
		&pk.Y,
		&pk.Z,
		&pk.Pitch,
		&pk.Yaw,
		&pk.VelocityX,
		&pk.VelocityY,
		&pk.VelocityZ,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
