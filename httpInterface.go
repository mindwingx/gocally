package gocally

type HttpInterface interface {
	SetUrl(url string) *HttpCall
	SetAuthorization(key string) *HttpCall
	SetHeader(key string, value string) *HttpCall
	SetHeaderBulk(headers map[string]string) *HttpCall
	DisableJsonHeaders() *HttpCall
	SetQueryParam(key string, value string) *HttpCall
	SetQueryParamBulk(query map[string]string) *HttpCall
	SetBody(payload any) *HttpCall
	SetForm(key string, value string) *HttpCall
	SetFormBulk(payload map[string]string) *HttpCall
	SetRequestTimeout(timeout int) *HttpCall
	Get() (map[string]interface{}, error)
	Post() (map[string]interface{}, error)
	Put() (map[string]interface{}, error)
	Delete() (map[string]interface{}, error)
}
