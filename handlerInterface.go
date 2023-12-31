package gocally

import "net/http"

type HandlerInterface interface {
	Response() (*http.Response, error)
	Payload() (map[string]interface{}, error)
	Entity(interface{}) (map[string]interface{}, error)
}
