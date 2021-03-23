package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//Initialisation
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("ContentType", "application/json")
	commonHeaders.Set("User-Agent", "Cool-http-client")
	client.builder.headers = commonHeaders

	//Execution
	requestHeader := make(http.Header)
	requestHeader.Set("X-request-Id", "ABC-123")

	finalHeaders := client.getRequestHEaders(requestHeader)

	//validation
	if len(finalHeaders) != 3 {
		t.Error("except 3 header")
	}
	if finalHeaders.Get("X-request-Id") != "ABC-123" {
		t.Error("invalid id recived")
	}

	if finalHeaders.Get("User-Agent") != "Cool-http-client" {
		t.Error("invalid id recived")
	}
	if finalHeaders.Get("ContentType") != "application/json" {
		t.Error("invalid id recived")
	}

}

func TestGetRequestBody(t *testing.T) {

	client := httpClient{}
	t.Run("noBodynilresponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Error("no error expected with nil body")
		}

		if body != nil {
			t.Error("no error expected with nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {})
	t.Run("BodyWithXmls", func(t *testing.T) {})
	t.Run("BBodyWithJsonAsDefault", func(t *testing.T) {})

}
