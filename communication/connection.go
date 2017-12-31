package communication

import (
	"bufio"
	"net"
)

type Connection struct {
	Conn        *net.Conn
	Writer      *bufio.Writer
	BufferWrite []byte
	Reader      *bufio.Reader
	BufferRead  []byte
	Key         string
}
