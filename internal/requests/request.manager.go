package requests

import (
	"fmt"
	"strings"
)

type RequestManager struct {
	requestStr  string
	RequestLine string
	Headers     []string
	Body        string
	R           Request
}

func NewRequestManager(rqStr string) (*RequestManager, error) {
	rqChunks, err := chunkifyHTTPRequestString(rqStr)
	if err != nil {
		return nil, fmt.Errorf(
			"could not split incoming request by `\r\n`: %v",
			err,
		)
	}

	reqLine := rqChunks[0]
	r, err := NewRequest(reqLine)
	if err != nil {
		return nil, fmt.Errorf(
			"could not create a new Request",
		)
	}

	// TODO assert here as in rqChunks should have rqChunks[0] etc
	return &RequestManager{
		requestStr:  rqStr,
		RequestLine: reqLine,
		Headers:     rqChunks[1:],
		Body:        "", // Also implement bodies
		R:           *r,
	}, nil
}

func chunkifyHTTPRequestString(rqStr string) ([]string, error) {
	return strings.Split(rqStr, "\r\n"), nil
}
