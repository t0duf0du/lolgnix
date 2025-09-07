package requests

import (
	"strings"
)

type ReqType string

const (
	GETREQTYPE ReqType = "GET"
)

type Request struct {
	reqType ReqType
	URLPath string
}

func NewRequest(reqLine string) (*Request, error) {
	reqLineChunks := strings.Split(reqLine, " ")
	return &Request{
		reqType: ReqType(reqLineChunks[0]),
		URLPath: reqLineChunks[1],
	}, nil
}
