package httpclient

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

// http_dto represents http request, http response field

// RequestArgs represents 3rd api necessary arguments
type RequestArgs struct {
	// RequestURL 3rd api url (required)
	RequestURL string
	// required
	Method string
	Body   []byte
	Params url.Values
	// Header http header in map format
	Header map[string]string
	// Timeout call timeout, request will cancel after this duration (if no retry)
	Timeout time.Duration
}

// validate required field
func (args RequestArgs) validate() error {
	if len(strings.TrimSpace(args.RequestURL)) == 0 {
		return errors.New("httpclient.RequestArgs.validate RequestURL is required")
	}
	if len(strings.TrimSpace(args.Method)) == 0 {
		return errors.New("httpclient.RequestArgs.validate Method is required")
	}

	return nil
}

// Response represents 3rd api response
type Response struct {
	Body []byte
	Code int
}
