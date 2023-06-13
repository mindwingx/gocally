package gocally

import "net/http"

type requestTransport struct {
	payload *http.Transport
	error   error
	next    RequestInterface
}

func (r *requestTransport) Process(item *RequestItems) {
	if r.error != nil {
		item.err = r.error
		return
	}

	if r.next != nil {
		r.next.Process(item)
	}
}

func (r *requestTransport) SetNext(next RequestInterface) {
	r.next = next
}
