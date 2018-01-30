package communication

import (
	"fmt"
	"github.com/pkg/errors"
)

type Request struct {
	Key      string
	Type     int
	Data     map[string]uint
	isSystem bool
}

func (r *Request) ValidadeFields(fields ...string) error {
	for _, field := range fields {
		if _, ok := r.Data[field]; !ok {
			return errors.New(fmt.Sprintf("bad %s value!", field))
		}

	}
	return nil
}
func (r *Request) ToAnswer() *Answer {
	return &Answer{0, r.Key, false, r.Type, "Bad Request", false}
}
