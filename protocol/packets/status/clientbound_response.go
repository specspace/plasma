package status

import (
	"github.com/specspace/plasma/protocol"
)

const ClientBoundResponsePacketID byte = 0x00

type ClientBoundResponse struct {
	JSONResponse protocol.String
}

func (pk ClientBoundResponse) Marshal() protocol.Packet {
	return protocol.MarshalPacket(
		ClientBoundResponsePacketID,
		pk.JSONResponse,
	)
}

func UnmarshalClientBoundResponse(packet protocol.Packet) (ClientBoundResponse, error) {
	var pk ClientBoundResponse

	if packet.ID != ClientBoundResponsePacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.JSONResponse,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
