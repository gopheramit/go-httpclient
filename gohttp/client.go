package gohttp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
	"sync"

	"github.com/gopheramit/go-httpclient/core"
)

type httpClient struct {
	builder *clientBuilder

	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*core.Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error)
	Delete(url string, headers ...http.Header) (*core.Response, error)
	Options(url string, headers ...http.Header) (*core.Response, error)
}

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
func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodGet, url, getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPut, url, getHeaders(headers...), nil)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodPatch, url, getHeaders(headers...), nil)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodDelete, url, getHeaders(headers...), nil)
}
func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.do(http.MethodOptions, url, getHeaders(headers...), nil)
}
