package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gopheramit/go-httpclient/gomime"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 2 * time.Second
	defaultConnectionTimeout  = 50 * time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)

	case gomime.ContentTypeXml:
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {

	fullHeaders := c.getRequestHeaders(headers)
	requestBody, err := c.getRequestBody(fullHeaders.Get(gomime.ContentType), body)
	if err != nil {
		return nil, errors.New("error creating request body")
	}

	if mock := mockUpServer.getMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("error creating request object")
	}

	request.Header = fullHeaders
	client := c.getHttpClient()

	resonse, err := client.Do(request)
	if err != nil {
		return nil, errors.New("error sending request")
	}
	responseBody, err := ioutil.ReadAll(resonse.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	defer resonse.Body.Close()
	finaleResponse := Response{
		status:     resonse.Status,
		statusCode: resonse.StatusCode,
		body:       responseBody,
		header:     resonse.Header,
	}

	return &finaleResponse, nil
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Timeout: c.getResponseTimeout() + c.getConnectionTimeout(),
			Transport: &http.Transport{
				MaxIdleConns:          c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}

	return defaultResponseTimeout
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}
