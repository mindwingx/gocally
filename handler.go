package gocally

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

//request

type RequestItems struct {
	url     requestUrl
	header  requestHeader
	query   requestQuery
	body    requestBody
	form    requestForm
	timeout requestTimeout
	method  string
	err     error
}

func PrepareRequestItems() *RequestItems {
	return &RequestItems{}
}

func (ri *RequestItems) SendRequest() (result HandlerInterface) {
	evaluateRequestItems(ri)
	if ri.err != nil {
		return newResponse(&http.Response{}, ri.err)
	}

	//init request
	timeout := ri.timeout.payload

	if timeout == 0 {
		timeout = 10
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	//set query params
	if len(ri.query.payload) > 0 {
		ri.url.payload += ri.query.payload
	}

	//prepare request instance
	req, err := http.NewRequest(ri.method, ri.url.payload, &ri.body.payload)

	if err != nil {
		return
	}

	//set request payload as the form value
	if len(ri.form.payload) > 0 {
		form := url.Values{}

		for formField, formValue := range ri.form.payload {
			form.Add(formField, formValue)
		}

		req.PostForm = form
	}

	//set headers
	if !ri.header.jsonHeaders {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	}

	for key, value := range ri.header.payload {
		req.Header.Set(key, value)
	}

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
	if r.response.Request != nil {
		defer r.response.Body.Close()
	}

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
