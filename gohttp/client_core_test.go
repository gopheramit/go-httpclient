package gohttp

import (
	"testing"
)

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
