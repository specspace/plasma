package status

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundRequestPacketID byte = 0x00

type ClientBoundRequest struct {
	JSONResponse protocol.String
}

func (pk ClientBoundRequest) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundRequestPacketID,
	)
}

func UnmarshalClientBoundRequest(packet protocol.Packet) (ClientBoundRequest, error) {
	var pk ClientBoundRequest

	if packet.ID != ClientBoundRequestPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	return pk, nil
}
