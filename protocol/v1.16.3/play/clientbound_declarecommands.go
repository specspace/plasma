package play

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundDeclareCommandsPacketID byte = 0x17

type ClientBoundDeclareCommands struct {
	Nodes     Nodes
	RootIndex packet.VarInt
}

type Node struct{}

type Nodes []Node

func (nodes Nodes) Encode() []byte {
	b := packet.VarInt(len(nodes)).Encode()
	return b
}

func (nodes *Nodes) Decode(r packet.DecodeReader) error {
	var length packet.VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	*nodes = make([]Node, length)
	return nil
}

func (pk ClientBoundDeclareCommands) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundDeclareCommandsPacketID,
		pk.Nodes,
		pk.RootIndex,
	)
}

func UnmarshalClientBoundDeclareCommands(packet packet.Packet) (ClientBoundDeclareCommands, error) {
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
