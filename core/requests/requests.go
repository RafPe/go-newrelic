package requests

// A Nrq is used to package data/params and operation
// used by client to send the request
type Nrq struct {
	Method       string
	Headers      map[string]string
	QueryParams  map[string]string
	RequestModel interface{}
}
