package gocally

import (
	"bytes"
)

type requestBody struct {
	payload bytes.Buffer
	error   error
	next    RequestInterface
}

func (r *requestBody) Process(item *RequestItems) {
	if r.error != nil {
		return
	}

	if r.next != nil {
		r.next.Process(item)
	}
}

func (r *requestBody) SetNext(next RequestInterface) {
	r.next = next
}
