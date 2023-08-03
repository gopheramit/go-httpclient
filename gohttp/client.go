package gohttp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
	"sync"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Put(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}

// func new() HttpClient {
// 	dialer := &net.Dialer{
// 		Timeout: 1 * time.Second,
// 	}
// 	client := http.Client{
// 		Transport: &http.Transport{
// 			MaxIdleConnsPerHost:   5,
// 			ResponseHeaderTimeout: 5 * time.Second,
// 			DialContext:           dialer.DialContext,
// 		},
// 	}
// 	httpClient := &httpClient{
// 		client: &client,
// 	}
// 	return httpClient
// }

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}
func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodGet, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
