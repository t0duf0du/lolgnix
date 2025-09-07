package requests

import "strings"

type ReqMethod string

const (
	GET ReqMethod = "GET"
)

type Request struct {
	Method ReqMethod
	Path   string
}

func NewRequest(requestLine string) (*Request, error) {
	chunksReqLine := strings.Split(requestLine, " ")
	return &Request{
		Method: ReqMethod(chunksReqLine[0]),
		Path:   chunksReqLine[1],
	}, nil
}
