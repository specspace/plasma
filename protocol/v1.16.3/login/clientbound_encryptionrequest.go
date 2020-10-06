package login

import (
	"github.com/spookspace/plasma/protocol"
	"github.com/spookspace/plasma/protocol/packet"
)

const ClientBoundEncryptionRequestPacketID byte = 0x01

type ClientBoundEncryptionRequest struct {
	ServerID    packet.String
	PublicKey   packet.ByteArray
	VerifyToken packet.ByteArray
}

func (pk ClientBoundEncryptionRequest) Marshal() packet.Packet {
	return packet.Marshal(
		ClientBoundEncryptionRequestPacketID,
		pk.ServerID,
		pk.PublicKey,
		pk.VerifyToken,
	)
}

func UnmarshalClientBoundEncryptionRequest(packet packet.Packet) (ClientBoundEncryptionRequest, error) {
	var pk ClientBoundEncryptionRequest

	if packet.ID != ClientBoundEncryptionRequestPacketID {
		return pk, protocol.ErrInvalidPacketID
	}

	if err := packet.Scan(
		&pk.ServerID,
		&pk.PublicKey,
		&pk.VerifyToken,
	); err != nil {
		return pk, err
	}

	return pk, nil
}
