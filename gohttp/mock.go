package gohttp

import (
	"fmt"
	"net/http"
)

// The mock structure provides clean way to configure http mocks
// based on combination betweeen method, url, request body, response body.
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	// fmt.Println("Mock response status code: ", m.ResponseStatusCode)
	// fmt.Println("Mock response body: ", m.ResponseBody)
	// fmt.Println("Mock response url: ", m.Url)
	response := Response{

		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}
	return &response, nil

}
