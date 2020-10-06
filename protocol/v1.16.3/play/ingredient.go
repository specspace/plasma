package play

import "github.com/spookspace/plasma/protocol/packet"

type Ingredient struct {
	Items Slots
}

func (ingredient Ingredient) Encode() []byte {
	return ingredient.Items.Encode()
}

func (ingredient *Ingredient) Decode(r packet.DecodeReader) error {
	return ingredient.Items.Decode(r)
}

type Ingredients []Ingredient

func (ingredients Ingredients) Encode() []byte {
	b := packet.VarInt(len(ingredients)).Encode()
	for _, ingredient := range ingredients {
		b = append(b, ingredient.Encode()...)
	}
	return b
}

func (ingredients *Ingredients) Decode(r packet.DecodeReader) error {
	var length packet.VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	*ingredients = make([]Ingredient, length)
	for _, ingredient := range *ingredients {
		if err := ingredient.Decode(r); err != nil {
			return err
		}
	}
	return nil
}
