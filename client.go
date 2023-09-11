package httpclient

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

// Client represents httpclient instance
type Client struct {
	fasthttpClient *fasthttp.Client
	logger         logger
}

type fasthttpConfigFunc func(*fasthttp.Client)

// NewClient create new Client instance with fasthttp.Client as dependency
// apply optional patterns to create with flexible setup
func NewClient(logger logger, options ...fasthttpConfigFunc) *Client {
	client := &Client{
		fasthttpClient: &fasthttp.Client{
			NoDefaultUserAgentHeader:  true,
			MaxIdemponentCallAttempts: 1,
		},
	}
	if logger == nil {
		logger = DefaultLogger()
	}
	client.logger = logger

	for _, opt := range options {
		opt(client.fasthttpClient)
	}

	return client
}

// Do http request with fasthttp lib
//
// params:
// ctx: context propagation
// args: requirement parameter for execute http request
//
// return:
// response, statusCode, and error (if has)
func (c *Client) Do(ctx context.Context, args *RequestArgs) (*Response, error) {
	var (
		// acquire request, response from pool
		req  = fasthttp.AcquireRequest()
		resp = fasthttp.AcquireResponse()
		err  error
	)

	// return req, resp object to pool
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	// validate argument
	err = args.validate()
	if err != nil {
		return &Response{
			Code: http.StatusBadRequest,
		}, err
	}

	// set up request
	requestURL := args.RequestURL
	if len(args.Params.Encode()) > 0 {
		requestURL += "?" + args.Params.Encode()
	}
	c.logger.Println(fmt.Sprintf("%s: %s", args.Method, requestURL))
	req.SetRequestURI(requestURL)
	req.Header.SetMethod(args.Method)

	if args.Body != nil {
		req.SetBody(args.Body)
	}
	if args.Header != nil {
		for k, v := range args.Header {
			req.Header.Add(k, v)
		}
	}

	if args.Timeout != 0 {
		err = c.fasthttpClient.DoTimeout(req, resp, args.Timeout)
	} else {
		err = c.fasthttpClient.Do(req, resp)
	}

	code := resp.StatusCode()
	if err != nil {
		return &Response{
			Code: code,
		}, err
	}
	if resp.StatusCode() >= http.StatusBadRequest {
		return &Response{
			Code: code,
		}, errors.New(string(resp.Body()))
	}

	return &Response{
		Body: resp.Body(),
		Code: code,
	}, nil
}

// WithMaxConnsPerHost set maximum number of connections per each host which may be established
func WithMaxConnsPerHost(maxConnsPerHost int) fasthttpConfigFunc {
	return func(c *fasthttp.Client) {
		// 512 is used if not set.
		// This mean only 512 http request can call at a time
		c.MaxConnsPerHost = maxConnsPerHost
	}
}

// WithIdleKeepAliveDuration set idle time for Keep-Alive connection
// Time reconnect = min(MaxIdleConnDuration,MaxConnDuration)
func WithIdleKeepAliveDuration(duration time.Duration) fasthttpConfigFunc {
	return func(c *fasthttp.Client) {
		// If an http keep alive instance do not exec in this duration
		// It will be kill
		// Default is 10s
		c.MaxIdleConnDuration = duration
	}
}

// WithMaxIdemponentCallAttempts set retry time when call api
// Default = 1
func WithMaxIdemponentCallAttempts(idemponent int) fasthttpConfigFunc {
	return func(c *fasthttp.Client) {
		// If a http call timeout is 1s
		// 3rd api reject http request after 0.1s
		// fasthttp has retry mechanism in this case
		// so it will retry (idempotent times, 1 by default)
		// result: call idempotent times
		c.MaxIdemponentCallAttempts = idemponent
	}
}

// WithDialTimeout tcp Dial with duration timeout
// If 3rd api down
// You make a http request to it without dialTimeout, it is very slow
// dialTimeout is different with readTimeout, writeTimeout
// dialTimeout is in TCP handshake phase
func WithDialTimeout(duration time.Duration) fasthttpConfigFunc {
	return func(c *fasthttp.Client) {
		c.Dial = func(addr string) (net.Conn, error) {
			// TODO optimize tcp dial
			conn, err := fasthttp.DialTimeout(addr, duration)
			if err != nil {
				return nil, fmt.Errorf("httpclient.dial %v", err)
			}

			return conn, nil
		}
	}
}

// WithReadBufferSize limit size of response
// This is also limits the maximum header size
// func WithReadBufferSize(readBufferSize int) fasthttpConfigFunc {
// 	return func(c *fasthttp.Client) {
// 		c.ReadBufferSize = readBufferSize
// 	}
// }

// WithMaxResponseBodySize limit response body size
func WithMaxResponseBodySize(reponseBodySize int) fasthttpConfigFunc {
	return func(c *fasthttp.Client) {
		c.MaxResponseBodySize = reponseBodySize
	}
}
