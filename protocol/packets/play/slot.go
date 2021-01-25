package play

import (
	"github.com/specspace/plasma/protocol"
)

type Slot struct {
	Present   protocol.Boolean
	ItemID    protocol.VarInt
	ItemCount protocol.Byte
	NBT       protocol.NBT
}

func (slot Slot) Encode() []byte {
	var b []byte
	b = append(b, slot.Present.Encode()...)
	if !slot.Present {
		return b
	}
	b = append(b, slot.ItemID.Encode()...)
	b = append(b, slot.ItemCount.Encode()...)
	b = append(b, slot.NBT.Encode()...)
	return b
}

func (slot *Slot) Decode(r protocol.DecodeReader) error {
	if err := slot.Present.Decode(r); err != nil {
		return err
	}
	if !slot.Present {
		return nil
	}
	if err := slot.ItemID.Decode(r); err != nil {
		return err
	}
	if err := slot.ItemCount.Decode(r); err != nil {
		return err
	}
	if err := slot.NBT.Decode(r); err != nil {
		return err
	}
	return nil
}

type Slots []Slot

func (slots Slots) Encode() []byte {
	b := protocol.VarInt(len(slots)).Encode()
	for _, slot := range slots {
		b = append(b, slot.Encode()...)
	}
	return b
}

func (slots *Slots) Decode(r protocol.DecodeReader) error {
	var length protocol.VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	*slots = make([]Slot, length)
	for _, slot := range *slots {
		if err := slot.Decode(r); err != nil {
			return err
		}
	}
	return nil
}
