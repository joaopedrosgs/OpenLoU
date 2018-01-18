package communication

type Request struct {
	Key      string
	Type     int
	Data     map[string]string
	isSystem bool
}

func (r *Request) ToAnswer() *Answer {
	return &Answer{0, r.Key, false, r.Type, "Bad Request", false}
}
