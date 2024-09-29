package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	maxIdleConnections int
	responseTimeout    time.Duration
	connectionTimeout  time.Duration
	disableTimeouts    bool
	headers            http.Header
}
type ClientBuilder interface {
	Build() Client
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(maxIdleConnections int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
}

func NewBuilder() ClientBuilder {
	httpClient := &clientBuilder{}
	return httpClient
}

func (c *clientBuilder) Build() Client {
	client := &httpClient{
		builder: c,
	}
	return client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(maxIdleConnections int) ClientBuilder {
	c.maxIdleConnections = maxIdleConnections
	return c
}
