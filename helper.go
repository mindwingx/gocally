package gocally

import "fmt"

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
