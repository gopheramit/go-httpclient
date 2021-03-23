package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"time"

	//"string"
	"strings"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

func (c *httpClient) getRequestBody(conentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(conentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {

	//client := http.Client{}
	fullHeaders := c.getRequestHEaders(headers)
	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("Unable to createnew request ")
	}

	request.Header = fullHeaders
	client := c.getHttpClient()

	return client.Do(request)
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}
	c.client = &http.Client{
		Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
			ResponseHeaderTimeout: c.getResponseTimeout(),
			DialContext: (&net.Dialer{
				Timeout: c.getConnectionTimeout(),
			}).DialContext,
		},
	}
	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {

	if c.responseTimeout > 0 {
		return c.responseTimeout
	}
	if c.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {

	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}
	if c.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getRequestHEaders(requestHeader http.Header) http.Header {
	result := make(http.Header)
	//Add common  headers from http client instance
	for header, value := range c.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	//Add custom headers from the  current request
	for header, value := range requestHeader {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}
