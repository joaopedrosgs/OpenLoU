package communication

import (
	"sync/atomic"
)

type Request struct {
	internal_id int32
	Key         string
	Type        int
	Data        map[string]string
	isSystem    bool
}

func (r *Request) ToAnswer() *Answer {

	return &Answer{r.internal_id, r.Key, r.Type, map[string]string{"Result": "False", "Message": "Bad request"}, false}
}

func (r *Request) SetInternalID(id int32) {
	atomic.StoreInt32(&r.internal_id, id)
}
func (request *Request) GetInternalID() int32 {
	return request.internal_id
}
