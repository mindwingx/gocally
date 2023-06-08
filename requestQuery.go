package gocally

type requestQuery struct {
	payload string
	error   error
	next    RequestInterface
}

func (r *requestQuery) Process(item *RequestItems) {
	if r.error != nil {
		item.err = r.error
		return
	}

	if r.next != nil {
		r.next.Process(item)
	}
}

func (r *requestQuery) SetNext(next RequestInterface) {
	r.next = next
}
