package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundTagsPacketID byte = 0x5c

type ClientBoundTags struct {
	BlockTags  Tags
	ItemTags   Tags
	FluidTags  Tags
	EntityTags Tags
}

type Tag struct {
	Type    packet.Identifier
	Entries packet.VarIntArray
}

type Tags []Tag

func (tags Tags) Encode() []byte {
	b := packet.VarInt(len(tags)).Encode()
	for _, tag := range tags {
		b = append(b, tag.Type.Encode()...)
		b = append(b, tag.Entries.Encode()...)
	}
	return b
}

func (tags *Tags) Decode(r packet.DecodeReader) error {
	var length packet.VarInt
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

func (pk ClientBoundTags) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundTagsPacketID,
		pk.BlockTags,
		pk.ItemTags,
		pk.FluidTags,
		pk.EntityTags,
	)
}

func UnmarshalClientBoundTags(packet packet.Packet) (ClientBoundTags, error) {
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
