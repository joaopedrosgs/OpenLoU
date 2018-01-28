package communication

import (
	"bufio"
	"net"
)

type User struct {
	Conn        *net.Conn
	Writer      *bufio.Writer
	BufferWrite []byte
	Reader      *bufio.Reader
	BufferRead  []byte
	Key         string
}
