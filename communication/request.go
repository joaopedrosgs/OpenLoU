package communication

import (
	"bufio"
	"net"
	"sync/atomic"
)

type Request struct {
	internal_id int32
	Key         string
	Type        int
	Data        map[string]string
	isSystem    bool
	conn        net.Conn
	Reader      *bufio.Reader
	BufferRead  []byte
}

func (r *Request) ToAnswer() *Answer {

	return &Answer{r.internal_id, r.Key, r.Type, false, "", false, r.conn, bufio.NewWriter(r.conn), make([]byte, 1024)}

}
func (r *Request) GetConn() net.Conn {
	return r.conn
}

func (r *Request) SetConn(conn net.Conn) {
	r.Reader = bufio.NewReader(conn)
	if r.BufferRead == nil {
		r.BufferRead = make([]byte, 1024)
	}
	r.conn = conn
}

func (r *Request) SetInternalID(id int32) {
	atomic.StoreInt32(&r.internal_id, id)
}
func (request *Request) GetInternalID() int32 {
	return request.internal_id
}
