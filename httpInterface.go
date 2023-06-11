package gocally

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
	SetRequestTimeout(int) *HttpCall
	Get() (map[string]interface{}, error)
	Post() (map[string]interface{}, error)
	Put() (map[string]interface{}, error)
	Delete() (map[string]interface{}, error)
}
