package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := httpClient{}
	commonHeader := make(http.Header)
	commonHeader.Set("Content-Type", "applicaiton/json")
	commonHeader.Set("User-Agent", "cool-http-client")
	client.builder.headers = commonHeader

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers ")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id recived")
	}
	if finalHeaders.Get("Content-Type") != "applicaiton/json" {
		t.Error("invalid  content type recived")
	}
	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent recived")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	t.Run("NoBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("no error expected while passing nil body")
		}
		if body != nil {
			t.Error("no body expected while passing nil body")
		}
	})
	t.Run("BodywithJson", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)
		if err != nil {
			t.Error("no error expected while marshaling a slice as json")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}

	})
	t.Run("BodywithXml", func(t *testing.T) {

	})
	t.Run("BodywithJsonDefault", func(t *testing.T) {

	})

}
