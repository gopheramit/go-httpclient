package gohttp

import (
	"net/http"

	"github.com/gopheramit/go-httpclient/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Add(header, value[0])
		}
	}

	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Add(header, value[0])
		}
	}
	//Set  user agenet if its not thier yet.
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) == "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}
