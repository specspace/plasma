package play

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundDeclareCommandsPacketID byte = 0x17

type ClientBoundDeclareCommands struct {
	Nodes     Nodes
	RootIndex protocol.VarInt
}

type Node struct{}

type Nodes []Node

func (nodes Nodes) Encode() []byte {
	b := protocol.VarInt(len(nodes)).Encode()
	return b
}

func (nodes *Nodes) Decode(r protocol.DecodeReader) error {
	var length protocol.VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	*nodes = make([]Node, length)
	return nil
}

func (pk ClientBoundDeclareCommands) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundDeclareCommandsPacketID,
		pk.Nodes,
		pk.RootIndex,
	)
}

func UnmarshalClientBoundDeclareCommands(packet protocol.Packet) (ClientBoundDeclareCommands, error) {
	var pk ClientBoundDeclareCommands

	if packet.ID != ClientBoundDeclareCommandsPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.Nodes,
		&pk.RootIndex,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
