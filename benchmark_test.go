package httpclient

import "testing"

// !!! run mock/main.go before execute benchmark
// go run mock/main.go

func Benchmark_callWithHTTPClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callWithHTTPClient()
	}
}

func Benchmark_callWithDefaultClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callWithDefaultClient()
	}
}
