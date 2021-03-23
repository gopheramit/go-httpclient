package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	client             *http.Client
	disableTimeouts    bool
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
}

type ClientBuilder interface {
	Build() Client
	SetHeaders(headers http.Header) ClientBuilder
	DisableTimeouts(b bool) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
}

func NewBuilder() ClientBuilder {

	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		headers:            c.headers,
		disableTimeouts:    c.disableTimeouts,
		maxIdleConnections: c.maxIdleConnections,
		connectionTimeout:  c.connectionTimeout,
		responseTimeout:    c.responseTimeout,
	}

	return &client
}

func (c *clientBuilder) DisableTimeouts(disableTimeouts bool) ClientBuilder {
	c.disableTimeouts = disableTimeouts
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
func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}
