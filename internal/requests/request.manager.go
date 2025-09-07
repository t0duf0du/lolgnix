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

	requestLine := rqChunks[0]
	header := rqChunks[1:]

	r, err := NewRequest(requestLine)
	if err != nil {
		return nil, fmt.Errorf(
			"could not create a new request from requerequestLine: %s. Got error: %v",
			requestLine,
			err,
		)
	}

	// TODO assert here as in rqChunks should have rqChunks[0] etc
	return &RequestManager{
		requestStr:  rqStr,
		RequestLine: requestLine,
		Headers:     header,
		Body:        "", // Also implement bodies
		R:           *r,
	}, nil
}

func chunkifyHTTPRequestString(rqStr string) ([]string, error) {
	return strings.Split(rqStr, "\r\n"), nil
}
