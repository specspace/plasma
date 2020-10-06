package packet

import (
	"bytes"
	"fmt"
	"github.com/spookspace/plasma/protocol/packet/zlib"
	"io"
)

// A Packet is the raw representation of message that is send between the client and the server
type Packet struct {
	ID   byte
	Data []byte
}

// Scan decodes and copies the Packet data into the fields
func (pk Packet) Scan(fields ...FieldDecoder) error {
	return Scan(bytes.NewReader(pk.Data), fields...)
}

// Marshal encodes the packet and compresses it if it is larger then the given threshold
func (pk *Packet) Marshal(threshold int) ([]byte, error) {
	var packedData []byte
	data := []byte{pk.ID}
	data = append(data, pk.Data...)

	if threshold > 0 {
		if len(data) > threshold {
			length := VarInt(len(data)).Encode()
			var err error
			data, err = zlib.Encode(data)
			if err != nil {
				return nil, err
			}

			packedData = append(packedData, VarInt(len(length)+len(data)).Encode()...)
			packedData = append(packedData, length...)
		} else {
			packedData = append(packedData, VarInt(int32(len(data)+1)).Encode()...)
			packedData = append(packedData, 0x00)
		}
	} else {
		packedData = append(packedData, VarInt(int32(len(data))).Encode()...)
	}

	return append(packedData, data...), nil
}

// Scan decodes a byte stream into fields
func Scan(r DecodeReader, fields ...FieldDecoder) error {
	for _, v := range fields {
		err := v.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

// Marshal transforms an ID and Fields into a Packet
func Marshal(ID byte, fields ...FieldEncoder) Packet {
	var pkt Packet
	pkt.ID = ID

	for _, v := range fields {
		pkt.Data = append(pkt.Data, v.Encode()...)
	}

	return pkt
}

// Unmarshal decodes and decompresses a byte array into a Packet
func Unmarshal(data []byte) (Packet, error) {
	reader := bytes.NewBuffer(data)

	var dataLength VarInt
	if err := dataLength.Decode(reader); err != nil {
		return Packet{}, err
	}

	decompressedData := make([]byte, dataLength)
	isCompressed := dataLength != 0
	if isCompressed {
		var err error
		if err = zlib.Decode(reader.Bytes(), decompressedData); err != nil {
			return Packet{}, err
		}
	} else {
		decompressedData = data[1:]
	}

	return Packet{
		ID:   decompressedData[0],
		Data: decompressedData[1:],
	}, nil
}

// ReadBytes decodes a byte stream and cuts the first Packet as a byte array out
func ReadBytes(r DecodeReader) ([]byte, error) {
	var packetLength VarInt
	if err := packetLength.Decode(r); err != nil {
		return nil, err
	}

	if packetLength < 1 {
		return nil, fmt.Errorf("packet length too short")
	}

	data := make([]byte, packetLength)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, fmt.Errorf("reading the content of the packet failed: %v", err)
	}

	return data, nil
}

// Read decodes and decompresses a byte stream and cuts the first Packet out
func Read(r DecodeReader, isZlib bool) (Packet, error) {
	data, err := ReadBytes(r)
	if err != nil {
		return Packet{}, err
	}

	if isZlib {
		return Unmarshal(data)
	}

	return Packet{
		ID:   data[0],
		Data: data[1:],
	}, nil
}

// Peek decodes and decompresses a byte stream and peeks the first Packet
func Peek(p PeekReader, isZlib bool) (Packet, error) {
	r := bytePeeker{
		PeekReader: p,
		cursor:     0,
	}

	return Read(&r, isZlib)
}
