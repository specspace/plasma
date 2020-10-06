package status

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundRequestPacketID byte = 0x00

type ClientBoundRequest struct {
	JSONResponse packet.String
}

func (pk ClientBoundRequest) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundRequestPacketID,
	)
}

func UnmarshalClientBoundRequest(packet packet.Packet) (ClientBoundRequest, error) {
	var pk ClientBoundRequest

	if packet.ID != ClientBoundRequestPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	return pk, nil
}
