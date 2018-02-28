package communication

import (
	"fmt"
)

type Request struct {
	Key      string
	Type     int
	Data     map[string]uint
	isSystem bool
}

func (r *Request) FieldsExist(fields ...string) error {
	for _, field := range fields {
		if _, ok := r.Data[field]; !ok {
			return fmt.Errorf("Empty %s value!", field)
		}

	}
	return nil
}
func (r *Request) ToAnswer() *Answer {
	return &Answer{0, r.Key, false, r.Type, "Bad Request", false}
}
