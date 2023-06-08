package gocally

type requestTimeout struct {
	payload int
	error   error
	next    RequestInterface
}

func (r *requestTimeout) Process(item *RequestItems) {
	if r.error != nil {
		item.err = r.error
		return
	}

	r.next.Process(item)
}

func (r *requestTimeout) SetNext(next RequestInterface) {
	r.next = next
}
