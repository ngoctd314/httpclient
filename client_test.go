package httpclient

import (
	"context"
	"net/url"
	"testing"
)

func TestClient_Do(t *testing.T) {
	client := NewClient(nil)
	params := url.Values{}
	params.Add("key", "value")
	client.Do(context.Background(), &RequestArgs{
		RequestURL: "test",
		Method:     "GET",
		Params:     params,
	})
}
