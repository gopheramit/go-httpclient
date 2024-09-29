package gohttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status     string
	statusCode int
	body       []byte
	header     http.Header
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) Headers() http.Header {
	return r.header
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) UnmarshalJSON(v interface{}) error {
	return json.Unmarshal(r.body, v)
}
