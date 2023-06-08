package gocally

type requestForm struct {
	payload map[string]string
	error   error
	next    RequestInterface
}

func (r *requestForm) Process(item *RequestItems) {
	if r.error != nil {
		item.err = r.error
		return
	}

	if r.next != nil {
		r.next.Process(item)
	}
}

func (r *requestForm) SetNext(next RequestInterface) {
	r.next = next
}
