package httpclient

import (
	"io"
	"net/http"
	"time"
)

var withHTTPClient = NewClient(
	nil,
	WithMaxConnsPerHost(4096),
	WithMaxIdemponentCallAttempts(1),
	WithIdleKeepAliveDuration(time.Second*10),
)
var withDefaultClient = http.Client{}

var (
	apiURL = "http://localhost:8080"
	method = http.MethodGet
)

func callWithHTTPClient() {
	// resp := withHTTPClient.Do(context.Background(), &RequestArgs{
	// 	RequestURL: apiURL,
	// 	Method:     method,
	// })
	// body is closed by fasthttp
	// _ = resp.Body
}

func callWithDefaultClient() {
	resp, _ := withDefaultClient.Get(apiURL)
	data, _ := io.ReadAll(resp.Body)
	_ = data
	// make sure close body
	resp.Body.Close()
}
