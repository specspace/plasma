package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundTagsPacketID byte = 0x5c

type ClientBoundTags struct {
	BlockTags  Tags
	ItemTags   Tags
	FluidTags  Tags
	EntityTags Tags
}

type Tag struct {
	Type    protocol.Identifier
	Entries protocol.VarIntArray
}

type Tags []Tag

func (tags Tags) Encode() []byte {
	b := protocol.VarInt(len(tags)).Encode()
	for _, tag := range tags {
		b = append(b, tag.Type.Encode()...)
		b = append(b, tag.Entries.Encode()...)
	}
	return b
}

func (tags *Tags) Decode(r protocol.DecodeReader) error {
	var length protocol.VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	*tags = make([]Tag, length)
	for _, tag := range *tags {
		if err := tag.Type.Decode(r); err != nil {
			return err
		}
		if err := tag.Entries.Decode(r); err != nil {
			return err
		}
	}
	return nil
}

func (pk ClientBoundTags) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundTagsPacketID,
		pk.BlockTags,
		pk.ItemTags,
		pk.FluidTags,
		pk.EntityTags,
	)
}

func UnmarshalClientBoundTags(packet protocol.Packet) (ClientBoundTags, error) {
	var pk ClientBoundTags

	if packet.ID != ClientBoundTagsPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.BlockTags,
		&pk.ItemTags,
		&pk.FluidTags,
		&pk.EntityTags,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
