package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ServerBoundEncryptionResponsePacketID = 0x01

type ServerBoundEncryptionResponse struct {
	SharedSecret packet.ByteArray
	VerifyToken  packet.ByteArray
}

func (pk ServerBoundEncryptionResponse) Marshal() packet.Packet {
	return packet.Marshal(
		ServerBoundEncryptionResponsePacketID,
		pk.SharedSecret,
		pk.VerifyToken,
	)
}

func UnmarshalServerBoundEncryptionResponse(packet packet.Packet) (ServerBoundEncryptionResponse, error) {
	var pk ServerBoundEncryptionResponse

	if packet.ID != ServerBoundEncryptionResponsePacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.SharedSecret,
		&pk.VerifyToken,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
