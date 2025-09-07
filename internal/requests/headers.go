package headers

import (
	"fmt"
	"strings"
)

type RequestManager struct {
	requestStr  string
	RequestLine string
	Headers     []string
	Body        string
}

func NewRequestManager(rqStr string) (*RequestManager, error) {
	rqChunks, err := chunkifyHTTPRequestString(rqStr)
	if err != nil {
		return nil, fmt.Errorf(
			"could not split incoming request by `\r\n`: %v",
			err,
		)
	}
	// TODO assert here as in rqChunks should have rqChunks[0] etc
	return &RequestManager{
		requestStr:  rqStr,
		RequestLine: rqChunks[0],
		Headers:     rqChunks[1:],
		Body:        "", // Also implement bodies
	}, nil
}

func chunkifyHTTPRequestString(rqStr string) ([]string, error) {
	return strings.Split(rqStr, "\r\n"), nil
}
