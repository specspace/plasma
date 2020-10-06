package status

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundResponsePacketID byte = 0x00

type ClientBoundResponse struct {
	JSONResponse packet.String
}

func (pk ClientBoundResponse) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundResponsePacketID,
		pk.JSONResponse,
	)
}

func UnmarshalClientBoundResponse(packet packet.Packet) (ClientBoundResponse, error) {
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
