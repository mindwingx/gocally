package gocally

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpCall struct {
	requestItems *RequestItems
	response     HandlerInterface
}

func SetRequest() *HttpCall {
	return &HttpCall{
		requestItems: PrepareRequestItems(),
	}
}

func (h *HttpCall) WithUrl(url string) *HttpCall {
	h.requestItems.url.payload = url
	return h
}

func (h *HttpCall) SetAuthorization(token string) *HttpCall {
	if len(h.requestItems.header.payload) == 0 {
		h.requestItems.header.payload = map[string]string{"Authorization": token}
	} else {
		h.requestItems.header.payload["Authorization"] = token
	}

	return h
}

func (h *HttpCall) SetHeader(key string, value string) *HttpCall {
	if len(h.requestItems.header.payload) == 0 {
		h.requestItems.header.payload = map[string]string{key: value}
	} else {
		h.requestItems.header.payload[key] = value
	}

	return h
}

func (h *HttpCall) SetHeaderBulk(headers map[string]string) *HttpCall {
	if len(h.requestItems.header.payload) == 0 {
		h.requestItems.header.payload = headers
	} else {
		for key, value := range headers {
			h.requestItems.header.payload[key] = value
		}
	}

	return h
}

func (h *HttpCall) DisableJsonHeaders() *HttpCall {
	h.requestItems.header.jsonHeaders = true
	return h
}

func (h *HttpCall) SetQueryParam(key string, value string) *HttpCall {
	h.requestItems.query.payload = QueryParamJoinSign(h.requestItems.query.payload)
	h.requestItems.query.payload += fmt.Sprintf("%s=%s", key, value)

	return h
}

func (h *HttpCall) SetQueryParamBulk(query map[string]string) *HttpCall {
	h.requestItems.query.payload = QueryParamJoinSign(h.requestItems.query.payload)
	setAmpersand, queryCount, loopCount := false, len(query), 0

	if queryCount > 1 {
		setAmpersand = true
	}

	for key, value := range query {
		loopCount++
		h.requestItems.query.payload += fmt.Sprintf("%s=%s", key, value)

		if setAmpersand && (loopCount < queryCount) {
			h.requestItems.query.payload += "&"
		}
	}

	return h
}

func (h *HttpCall) SetBody(payload any) *HttpCall {
	err := json.NewEncoder(&h.requestItems.body.payload).Encode(payload)

	if err != nil {
		h.requestItems.body.error = err
	}

	return h
}

func (h *HttpCall) SetForm(key string, value string) *HttpCall {
	if len(h.requestItems.form.payload) == 0 {
		h.requestItems.form.payload = map[string]string{key: value}
	} else {
		h.requestItems.form.payload[key] = value
	}

	return h
}

func (h *HttpCall) SetFormBulk(payload map[string]string) *HttpCall {
	if len(h.requestItems.form.payload) == 0 {
		h.requestItems.form.payload = payload
	} else {
		for key, value := range payload {
			h.requestItems.form.payload[key] = value
		}
	}

	return h
}

func (h *HttpCall) SetRequestTimeout(timeoutInSec int) *HttpCall {
	h.requestItems.timeout.payload = timeoutInSec
	return h
}

func (h *HttpCall) Get() HandlerInterface {
	h.requestItems.method = http.MethodGet
	return h.requestItems.SendRequest()
}

func (h *HttpCall) Post() HandlerInterface {
	h.requestItems.method = http.MethodPost
	return h.requestItems.SendRequest()
}

func (h *HttpCall) Put() HandlerInterface {
	h.requestItems.method = http.MethodPut
	return h.requestItems.SendRequest()
}

func (h *HttpCall) Delete() HandlerInterface {
	h.requestItems.method = http.MethodDelete
	return h.requestItems.SendRequest()
}
