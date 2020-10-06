package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const (
	ClientBoundJoinGamePacketID byte = 0x24
)

type ClientBoundJoinGame struct {
	EntityID            packet.VarInt
	IsHardcore          packet.Boolean
	Gamemode            packet.UnsignedByte
	PreviousGamemode    packet.UnsignedByte
	WorldNames          packet.IdentifierArray
	DimensionCodec      packet.NBT
	Dimension           packet.NBT
	WorldName           packet.Identifier
	HashedSeed          packet.Long
	MaxPlayers          packet.VarInt
	ViewDistance        packet.VarInt
	ReducedDebugInfo    packet.Boolean
	EnableRespawnScreen packet.Boolean
	IsDebug             packet.Boolean
	IsFlat              packet.Boolean
}

type DimensionCodecVanilla struct {
	Name               string  `nbt:"name"`
	PiglinSafe         byte    `nbt:"piglin_safe"`
	Natural            byte    `nbt:"natural"`
	AmbientLight       float32 `nbt:"ambient_light"`
	FixedTime          int     `nbt:"fixed_time"`
	Infiniburn         string  `nbt:"infiniburn"`
	RespawnAnchorWorks byte    `nbt:"respawn_anchor_works"`
	HasSkylight        byte    `nbt:"has_skylight"`
	BedWorks           byte    `nbt:"bed_works"`
	Effects            string  `nbt:"effects"`
	HasRaids           byte    `nbt:"has_raids"`
	LogicalHeight      int     `nbt:"logical_height"`
	CoordinateScale    float32 `nbt:"coordinate_scale"`
	Ultrawarm          byte    `nbt:"ultrawarm"`
	HasCeiling         byte    `nbt:"has_ceiling"`
}

func (pk ClientBoundJoinGame) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundJoinGamePacketID,
		pk.EntityID,
		pk.IsHardcore,
		pk.Gamemode,
		pk.PreviousGamemode,
		pk.WorldNames,
		pk.DimensionCodec,
		pk.Dimension,
		pk.WorldName,
		pk.HashedSeed,
		pk.MaxPlayers,
		pk.ViewDistance,
		pk.ReducedDebugInfo,
		pk.EnableRespawnScreen,
		pk.IsDebug,
		pk.IsFlat,
	)
}

func UnmarshalClientBoundJoinGame(packet packet.Packet) (ClientBoundJoinGame, error) {
	var pk ClientBoundJoinGame

	if packet.ID != ClientBoundJoinGamePacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.EntityID,
		&pk.IsHardcore,
		&pk.Gamemode,
		&pk.PreviousGamemode,
		&pk.WorldNames,
		&pk.DimensionCodec,
		&pk.Dimension,
		&pk.WorldName,
		&pk.HashedSeed,
		&pk.MaxPlayers,
		&pk.ViewDistance,
		&pk.ReducedDebugInfo,
		&pk.EnableRespawnScreen,
		&pk.IsDebug,
		&pk.IsFlat,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
