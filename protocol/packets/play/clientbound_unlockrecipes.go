package play

import "github.com/specspace/plasma/protocol"

const ClientBoundUnlockRecipesPacketID byte = 0x35

type ClientBoundUnlockRecipes struct {
	Action                             protocol.VarInt
	CraftingRecipeBookFilterActive     protocol.Boolean
	CraftingRecipeBookOpen             protocol.Boolean
	SmeltingRecipeBookFilterActive     protocol.Boolean
	SmeltingRecipeBookOpen             protocol.Boolean
	BlastFurnaceRecipeBookFilterActive protocol.Boolean
	BlastFurnaceRecipeBookOpen         protocol.Boolean
	SmokerRecipeBookFilterActive       protocol.Boolean
	SmokerRecipeBookOpen               protocol.Boolean
	RecipeIDs1                         protocol.IdentifierArray
	RecipeIDs2                         protocol.IdentifierArray
}

func (pk ClientBoundUnlockRecipes) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundUnlockRecipesPacketID,
		pk.Action,
		pk.CraftingRecipeBookFilterActive,
		pk.CraftingRecipeBookOpen,
		pk.SmeltingRecipeBookFilterActive,
		pk.SmeltingRecipeBookOpen,
		pk.BlastFurnaceRecipeBookFilterActive,
		pk.BlastFurnaceRecipeBookOpen,
		pk.SmokerRecipeBookFilterActive,
		pk.SmokerRecipeBookOpen,
		pk.RecipeIDs1,
		pk.RecipeIDs2,
	)
}

func UnmarshalClientBoundUnlockRecipes(packet protocol.Packet) (ClientBoundUnlockRecipes, error) {
	var pk ClientBoundUnlockRecipes

	if packet.ID != ClientBoundUnlockRecipesPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Action,
		&pk.CraftingRecipeBookFilterActive,
		&pk.CraftingRecipeBookOpen,
		&pk.SmeltingRecipeBookFilterActive,
		&pk.SmeltingRecipeBookOpen,
		&pk.BlastFurnaceRecipeBookFilterActive,
		&pk.BlastFurnaceRecipeBookOpen,
		&pk.SmokerRecipeBookFilterActive,
		&pk.SmokerRecipeBookOpen,
		&pk.RecipeIDs1,
		&pk.RecipeIDs2,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
