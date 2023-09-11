// Package httpclient is used as client to call 3rd api
// httpclient is based on https://github.com/valyala/fasthttp
// httpclient provide:
// + standard request/response format
// + standard http config (e.g http idle, http keep alive, timeout, read timeout, write timeout, retry, connection per host)
// + integrate with circuit breaker (using proxy pattern)
// +
package httpclient
