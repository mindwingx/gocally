package gocally

type requestHeader struct {
	payload     map[string]string
	jsonHeaders bool
	error       error
	next        RequestInterface
}

func (r *requestHeader) Process(item *RequestItems) {
	if r.error != nil {
		item.err = r.error
		return
	}

	if r.next != nil {
		r.next.Process(item)
	}
}

func (r *requestHeader) SetNext(next RequestInterface) {
	r.next = next
}
