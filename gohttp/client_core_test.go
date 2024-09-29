package gohttp

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("content-type", "application/json")
	commonHeaders.Set("User-Agent", "test-client")
	client.builder.headers = commonHeaders
	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "abc-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if finalHeaders.Get("content-type") != "application/json" {
		t.Error("invalid content-type header value")
	}
	if finalHeaders.Get("User-Agent") != "test-client" {
		t.Error("invalid User-Agent header value")
	}
	if finalHeaders.Get("X-Request-Id") != "abc-123" {
		t.Error("invalid X-Request-Id header value")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	t.Run("nil body", func(t *testing.T) {
		body, err := client.getRequestBody("application/json", nil)
		if err != nil {
			t.Error("we were not expecting an error")
		}

		if body != nil {
			t.Error("we were expecting a nil body")
		}
	})

	t.Run("json body", func(t *testing.T) {
		body := map[string]string{
			"hello": "world",
		}
		bodyBytes, err := client.getRequestBody("application/json", body)
		if err != nil {
			t.Error("we were not expecting an error")
		}

		if string(bodyBytes) != `{"hello":"world"}` {
			t.Error("invalid json body")
		}
	})

	t.Run("xml body", func(t *testing.T) {
		type Response struct {
			XMLName xml.Name `xml:"map"`
			Hello   string   `xml:"hello"`
		}
		body := Response{Hello: "world"}
		bodyBytes, err := client.getRequestBody("application/xml", body)
		if err != nil {
			t.Error("we were not expecting an error")
		}
		fmt.Println("xml response", string(bodyBytes))
		fmt.Println("error", err)
		if string(bodyBytes) != `<map><hello>world</hello></map>` {
			t.Error("invalid xml body")
		}
	})

	t.Run("unsupported body", func(t *testing.T) {
		body := map[string]string{
			"hello": "world",
		}
		bodyBytes, err := client.getRequestBody("application/unsupported", body)
		if err != nil {
			t.Error("we were not expecting an error")
		}

		if string(bodyBytes) != `{"hello":"world"}` {
			t.Error("invalid json body")
		}
	})
}
