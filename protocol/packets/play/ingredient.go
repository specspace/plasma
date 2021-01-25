package play

import (
	"errors"
	"github.com/specspace/plasma/protocol"
)

type Ingredient struct {
	Items Slots
}

func (ingredient Ingredient) Encode() []byte {
	return ingredient.Items.Encode()
}

func (ingredient *Ingredient) Decode(r protocol.DecodeReader) error {
	return ingredient.Items.Decode(r)
}

type Ingredients []Ingredient

func (ingredients Ingredients) Encode() []byte {
	b := protocol.VarInt(len(ingredients)).Encode()
	for _, ingredient := range ingredients {
		b = append(b, ingredient.Encode()...)
	}
	return b
}

func (ingredients *Ingredients) Decode(r protocol.DecodeReader) error {
	var length protocol.VarInt
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

type IngredientsShaped Ingredients

func (ingredientsShaped IngredientsShaped) Encode() []byte {
	var b []byte
	for _, ingredient := range ingredientsShaped {
		b = append(b, ingredient.Encode()...)
	}
	return b
}

func (ingredientsShaped *IngredientsShaped) Decode(r protocol.DecodeReader) error {
	if ingredientsShaped == nil {
		return errors.New("shaped ingredient cannot be null")
	}

	for _, ingredient := range *ingredientsShaped {
		if err := ingredient.Decode(r); err != nil {
			return err
		}
	}
	return nil
}
