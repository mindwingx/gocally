package gocally

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func QueryParamJoinSign(query string) (sign string) {
	sign = "?"

	if len(query) > 0 {
		sign = fmt.Sprintf("%s%s", query, "&")
	}

	return
}

func evaluateRequestItems(ri *RequestItems) {
	reqUrl := &ri.url

	reqHeader := &ri.header
	reqHeader.SetNext(reqUrl)

	reqQuery := &ri.query
	reqQuery.SetNext(reqHeader)

	reqBody := &ri.body
	reqBody.SetNext(reqQuery)

	reqForm := &ri.form
	reqForm.SetNext(reqBody)

	reqForm.Process(ri)
}

func handleHttpClient(ri *RequestItems, client *http.Client) {
	timeout := ri.timeout.payload

	if timeout == 0 {
		timeout = 5
	}

	client.Timeout = time.Second * time.Duration(timeout)

	if ri.transport.payload != nil {
		client.Transport = ri.transport.payload
	}
}

func handleHeaders(ri *RequestItems, req *http.Request) {
	if !ri.header.jsonHeaders {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
	}

	for key, value := range ri.header.payload {
		req.Header.Set(key, value)
	}
}

func handlePostForm(ri *RequestItems, req *http.Request) {
	if len(ri.form.payload) > 0 {
		form := url.Values{}

		for formField, formValue := range ri.form.payload {
			form.Add(formField, formValue)
		}

		req.PostForm = form
	}
}
