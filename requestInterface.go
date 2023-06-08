package gocally

type RequestInterface interface {
	Process(item *RequestItems)
	SetNext(next RequestInterface)
}
