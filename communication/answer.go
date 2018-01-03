package communication

import (
	"bufio"
	"net"
)

type Answer struct {
	internal_id int32
	key         string
	Type        int
	Result      bool
	Data        string
	isSystem    bool
	conn        net.Conn
	writer      *bufio.Writer
	bufferWrite []byte
}

func (answer *Answer) GetKey() string {
	return answer.key
}

var badRequest = Answer{0, "", -1, false, "Bad request", false, nil, nil, nil}
var unauthorizedRequest = Answer{0, "", -1, false, "Unauthorized", false, nil, nil, nil}

func BadRequest(conn net.Conn) *Answer {
	badRequest.SetConn(conn)
	return &badRequest
}
func Unauthorized(conn net.Conn) *Answer {
	unauthorizedRequest.SetConn(conn)
	return &unauthorizedRequest
}

func (answer *Answer) IsSystem() bool {
	return answer.isSystem
}

func (answer *Answer) GetConn() net.Conn {
	return answer.conn
}

func (answer *Answer) SetConn(conn net.Conn) {
	answer.writer = bufio.NewWriter(conn)
	answer.conn = conn
}

func (answer *Answer) GetWriter() *bufio.Writer {

	return answer.writer
}

func (answer *Answer) GetBuffer() *[]byte {
	return &answer.bufferWrite
}
func (answer *Answer) GetInternalID() int32 {
	return answer.internal_id
}
