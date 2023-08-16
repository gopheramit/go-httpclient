package gohttp_mock

import (
	"fmt"
	"net/http"

	"github.com/gopheramit/go-httpclient/core"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseStatusCode int
	ResponseBody       string
	ResponseHeaders    http.Header
}

func (m *Mock) GetResponse() (*core.Response, error) {
	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
		Headers:    make(http.Header),
	}

	for header, _ := range m.ResponseHeaders {
		response.Headers.Set(header, m.ResponseHeaders.Get(header))

	}
	return &response, nil
}
