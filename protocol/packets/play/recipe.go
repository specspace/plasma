package play

import (
	"github.com/specspace/plasma/protocol"
)

const (
	RecipeTypeCraftingShapeless               RecipeType = "crafting_shapeless"
	RecipeTypeCraftingShaped                  RecipeType = "crafting_shaped"
	RecipeTypeCraftingSpecialArmorDye         RecipeType = "crafting_special_armordye"
	RecipeTypeCraftingSpecialBookCloning      RecipeType = "crafting_special_bookcloning"
	RecipeTypeCraftingSpecialMapCloning       RecipeType = "crafting_special_mapcloning"
	RecipeTypeCraftingSpecialMapExtending     RecipeType = "crafting_special_mapextending"
	RecipeTypeCraftingSpecialFireworkRocket   RecipeType = "crafting_special_firework_rocket"
	RecipeTypeCraftingSpecialFireworkStar     RecipeType = "crafting_special_firework_star"
	RecipeTypeCraftingSpecialFireworkStarFade RecipeType = "crafting_special_firework_star_fade"
	RecipeTypeCraftingSpecialRepairItem       RecipeType = "crafting_special_repairitem"
	RecipeTypeCraftingSpecialTippedArrow      RecipeType = "crafting_special_tippedarrow"
	RecipeTypeCraftingSpecialBannerDuplicate  RecipeType = "crafting_special_bannerduplicate"
	RecipeTypeCraftingSpecialBannerAddPattern RecipeType = "crafting_special_banneraddpattern"
	RecipeTypeCraftingSpecialShieldDecoration RecipeType = "crafting_special_shielddecoration"
	RecipeTypeCraftingSpecialShulkerBoxColor  RecipeType = "crafting_special_shulkerboxcoloring"
	RecipeTypeCraftingSpecialSuspiciousStew   RecipeType = "crafting_special_suspiciousstew"
	RecipeTypeSmelting                        RecipeType = "smelting"
	RecipeTypeBlasting                        RecipeType = "blasting"
	RecipeTypeSmoking                         RecipeType = "smoking"
	RecipeTypeCampfireCooking                 RecipeType = "campfire_cooking"
	RecipeTypeStoneCutting                    RecipeType = "stonecutting"
)

type RecipeCraftingShaped struct {
	Width       protocol.VarInt
	Height      protocol.VarInt
	Group       protocol.String
	Ingredients IngredientsShaped
	Result      Slot
}

func (recipe RecipeCraftingShaped) Encode() []byte {
	recipe.Ingredients = make([]Ingredient, recipe.Width*recipe.Height)
	return protocol.MarshalFields(
		recipe.Width,
		recipe.Height,
		recipe.Group,
		recipe.Ingredients,
		recipe.Result,
	)
}

func (recipe *RecipeCraftingShaped) Decode(r protocol.DecodeReader) error {
	return protocol.ScanFields(r,
		&recipe.Width,
		&recipe.Height,
		&recipe.Group,
		&recipe.Ingredients,
		&recipe.Result,
	)
}

type RecipeCraftingShapeless struct {
	Group       protocol.String
	Ingredients Ingredients
	Result      Slot
}

func (recipe RecipeCraftingShapeless) Encode() []byte {
	return protocol.MarshalFields(
		recipe.Group,
		recipe.Ingredients,
		recipe.Result,
	)
}

func (recipe *RecipeCraftingShapeless) Decode(r protocol.DecodeReader) error {
	return protocol.ScanFields(r,
		&recipe.Group,
		&recipe.Ingredients,
		&recipe.Result,
	)
}

type RecipeCooking struct {
	Group       protocol.String
	Ingredient  Ingredient
	Result      Slot
	Experience  protocol.Float
	CookingTime protocol.VarInt
}

func (recipe RecipeCooking) Encode() []byte {
	return protocol.MarshalFields(
		recipe.Group,
		recipe.Ingredient,
		recipe.Result,
		recipe.Experience,
		recipe.CookingTime,
	)
}

func (recipe *RecipeCooking) Decode(r protocol.DecodeReader) error {
	return protocol.ScanFields(r,
		&recipe.Group,
		&recipe.Ingredient,
		&recipe.Result,
		&recipe.Experience,
		&recipe.CookingTime,
	)
}

type RecipeStoneCutting struct {
	Group      protocol.String
	Ingredient Ingredient
	Result     Slot
}

func (recipe RecipeStoneCutting) Encode() []byte {
	return protocol.MarshalFields(
		recipe.Group,
		recipe.Ingredient,
		recipe.Result,
	)
}

func (recipe *RecipeStoneCutting) Decode(r protocol.DecodeReader) error {
	return protocol.ScanFields(r,
		&recipe.Group,
		&recipe.Ingredient,
		&recipe.Result,
	)
}

type RecipeType protocol.Identifier

func (recipeType RecipeType) Encode() []byte {
	return recipeType.Encode()
}

func (recipeType *RecipeType) Decode(r protocol.DecodeReader) error {
	return recipeType.Decode(r)
}

type Recipe struct {
	Type     RecipeType
	RecipeID protocol.String
	Data     protocol.Field
}

func (recipe Recipe) Encode() []byte {
	return protocol.MarshalFields(
		recipe.Type,
		recipe.RecipeID,
		recipe.Data,
	)
}

func (recipe *Recipe) Decode(r protocol.DecodeReader) error {
	switch recipe.Type {
	/*case RecipeTypeCraftingShaped:
	data, ok := recipe.Data.(RecipeCraftingShaped)
	if !ok
		b = append(b, data.Encode()...)*/
	}

	return protocol.ScanFields(r,
		&recipe.Type,
		&recipe.RecipeID,
		//&recipe.Data,
	)
}

/*func (recipe Recipe) DataAsCraftingShapeless() (RecipeCraftingShapeless, error) {
	ok, craftingShapeless := recipe.Data.(RecipeCraftingShapeless)
	if !ok {
		return RecipeCraftingShapeless{}, errors.New("")
	}

	return craftingShapeless, nil
}*/

type Recipes []Recipe

func (recipes Recipes) Encode() []byte {
	b := protocol.VarInt(len(recipes)).Encode()
	for _, recipe := range recipes {
		b = append(b, recipe.Encode()...)
	}
	return b
}

func (recipes *Recipes) Decode(r protocol.DecodeReader) error {
	var length protocol.VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	*recipes = make([]Recipe, length)
	for _, recipe := range *recipes {
		if err := recipe.Decode(r); err != nil {
			return err
		}
	}
	return nil
}
