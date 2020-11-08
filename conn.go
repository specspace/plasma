package plasma

import (
	"bufio"
	"crypto/cipher"
	"github.com/specspace/plasma/protocol"
	"io"
	"net"
)

type PacketWriter interface {
	WritePacket(p protocol.Packet) error
}

type PacketReader interface {
	ReadPacket() (protocol.Packet, error)
}

type conn struct {
	net.Conn

	r         *bufio.Reader
	w         io.Writer
	state     protocol.State
	threshold int
}

// Conn is a minecraft Connection
type Conn interface {
	net.Conn
	PacketWriter
	PacketReader
	State() protocol.State
	Threshold() int
}

// wrapConn warp an net.Conn to plasma.Conn
func wrapConn(c net.Conn) conn {
	return conn{
		Conn:      c,
		r:         bufio.NewReader(c),
		w:         c,
		state:     protocol.StateHandshaking,
		threshold: 0,
	}
}

func (c conn) State() protocol.State {
	return c.state
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
