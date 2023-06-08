package gocally

import "errors"

type requestUrl struct {
	payload string
	error   error
	next    RequestInterface
}

func (r *requestUrl) Process(item *RequestItems) {
	if len(r.payload) == 0 {
		item.err = errors.New("error: no URL is set")
		return
	}

	if r.error != nil {
		item.err = r.error
		return
	}

	if r.next != nil {
		r.next.Process(item)
	}
}

func (r *requestUrl) SetNext(next RequestInterface) {
	r.next = next
}
