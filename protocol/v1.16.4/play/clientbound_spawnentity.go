package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundSpawnEntityPacketID byte = 0x00

type ClientBoundSpawnEntity struct {
	EntityID   protocol.VarInt
	ObjectUUID protocol.UUID
	Type       protocol.VarInt
	X          protocol.Double
	Y          protocol.Double
	Z          protocol.Double
	Pitch      protocol.Angle
	Yaw        protocol.Angle
	VelocityX  protocol.Short
	VelocityY  protocol.Short
	VelocityZ  protocol.Short
}

func (pk ClientBoundSpawnEntity) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
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

func UnmarshalClientBoundSpawnEntity(packet protocol.Packet) (ClientBoundSpawnEntity, error) {
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
