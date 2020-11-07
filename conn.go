package plasma

import (
	"bufio"
	"crypto/cipher"
	"github.com/specspace/plasma/protocol"
	"io"
	"net"
)

type conn struct {
	net.Conn

	r         *bufio.Reader
	w         io.Writer
	threshold int
}

// Conn is a minecraft Connection
type Conn interface {
	net.Conn
	Threshold() int
}

// WrapConn warp an net.Conn to plasma.Conn
func WrapConn(c net.Conn) Conn {
	return &conn{
		Conn:      c,
		r:         bufio.NewReader(c),
		w:         c,
		threshold: 0,
	}
}

// Threshold returns the number of bytes a Packet can be long before it will be compressed
func (c conn) Threshold() int {
	return c.threshold
}

func (c *conn) Read(b []byte) (int, error) {
	return c.r.Read(b)
}

func (c *conn) Write(b []byte) (int, error) {
	return c.w.Write(b)
}

// ReadPacket read a Packet from Conn.
func (c *conn) ReadPacket() (protocol.Packet, error) {
	return protocol.ReadPacket(c.r, c.threshold > 0)
}

// PeekPacket peeks a Packet from Conn.
func (c *conn) PeekPacket() (protocol.Packet, error) {
	return protocol.PeekPacket(c.r, c.threshold > 0)
}

//WritePacket write a Packet to Conn.
func (c *conn) WritePacket(p protocol.Packet) error {
	pk, err := p.Marshal(c.threshold)
	if err != nil {
		return err
	}
	_, err = c.w.Write(pk)
	return err
}

// SetCipher sets the decode/encode stream for this Conn
func (c *conn) SetCipher(ecoStream, decoStream cipher.Stream) {
	c.r = bufio.NewReader(cipher.StreamReader{
		S: decoStream,
		R: c.Conn,
	})
	c.w = cipher.StreamWriter{
		S: ecoStream,
		W: c.Conn,
	}
}
