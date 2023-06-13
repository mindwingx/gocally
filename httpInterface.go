package gocally

import "net/http"

type HttpInterface interface {
	WithUrl(string) *HttpCall
	SetAuthorization(string) *HttpCall
	SetHeader(string, string) *HttpCall
	SetHeaderBulk(map[string]string) *HttpCall
	DisableJsonHeaders() *HttpCall
	SetQueryParam(string, string) *HttpCall
	SetQueryParamBulk(map[string]string) *HttpCall
	SetBody(any) *HttpCall
	SetForm(string, string) *HttpCall
	SetFormBulk(map[string]string) *HttpCall
	SetHttpTimeout(int) *HttpCall
	SetHttpTransport(*http.Transport) *HttpCall
	Get() HandlerInterface
	Post() HandlerInterface
	Put() HandlerInterface
	Delete() HandlerInterface
}
