package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundDeclareRecipesPacketID byte = 0x5b

type ClientBoundDeclareRecipes struct {
	Recipes Recipes
}

func (pk ClientBoundDeclareRecipes) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundDeclareRecipesPacketID,
		pk.Recipes,
	)
}

func UnmarshalClientBoundDeclareRecipes(packet protocol.Packet) (ClientBoundDeclareRecipes, error) {
	var pk ClientBoundDeclareRecipes

	if packet.ID != ClientBoundDeclareRecipesPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Recipes,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
