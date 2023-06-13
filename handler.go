package gocally

import (
	"encoding/json"
	"net/http"
)

//request

type RequestItems struct {
	url       requestUrl
	header    requestHeader
	query     requestQuery
	body      requestBody
	form      requestForm
	timeout   requestTimeout
	transport requestTransport
	method    string
	err       error
}

func PrepareRequestItems() *RequestItems {
	return &RequestItems{}
}

func (ri *RequestItems) SendRequest() (result HandlerInterface) {
	evaluateRequestItems(ri)

	if ri.err != nil {
		return newResponse(&http.Response{}, ri.err)
	}

	client := &http.Client{}
	handleHttpClient(ri, client)

	if len(ri.query.payload) > 0 {
		ri.url.payload += ri.query.payload
	}

	req, err := http.NewRequest(ri.method, ri.url.payload, &ri.body.payload)

	if err != nil {
		return
	}

	handlePostForm(ri, req)
	handleHeaders(ri, req)

	response, err := client.Do(req)

	return newResponse(response, err)
}

// response

type Response struct {
	response *http.Response
	err      error
}

func newResponse(response *http.Response, err error) *Response {
	return &Response{response: response, err: err}
}

func (r *Response) Response() (*http.Response, error) {
	return r.response, r.err
}

func (r *Response) Payload() (payload map[string]interface{}, err error) {
	if r.response.Request != nil {
		defer r.response.Body.Close()
	}

	var resResult map[string]interface{}

	err = json.NewDecoder(r.response.Body).Decode(&resResult)

	if err != nil {
		return
	}

	payload = map[string]interface{}{
		"status":      r.response.Status,
		"status_code": r.response.StatusCode,
		"payload":     resResult,
	}

	return
}

func (r *Response) Entity(instance interface{}) (payload map[string]interface{}, err error) {
	if r.response.Request != nil {
		defer r.response.Body.Close()
	}

	err = json.NewDecoder(r.response.Body).Decode(instance)

	if err != nil {
		return
	}

	payload = map[string]interface{}{
		"status":      r.response.Status,
		"status_code": r.response.StatusCode,
	}

	return
}
