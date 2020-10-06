package zlib

import (
	"github.com/4kills/go-libdeflate"
	"sync"
)

var (
	decoder    libdeflate.Decompressor
	decodeLock sync.Mutex

	encoder    libdeflate.Compressor
	encodeLock sync.Mutex
)

func init() {
	var err error
	decoder, err = libdeflate.NewDecompressor()
	if err != nil {
		panic(err)
	}

	encoder, err = libdeflate.NewCompressor()
	if err != nil {
		panic(err)
	}
}

func Decode(in, out []byte) error {
	decodeLock.Lock()
	defer decodeLock.Unlock()

	_, _, err := decoder.DecompressZlib(in, out)
	return err
}

func Encode(in []byte) ([]byte, error) {
	encodeLock.Lock()
	defer encodeLock.Unlock()

	_, bb, err := encoder.CompressZlib(in, nil)
	return bb, err
}
