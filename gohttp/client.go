package gohttp

type httpClient struct{}

type HttpClient interface {
	Get()
	Put()
	Patch()
	Post()
	Delete()
}

func (c *httpClient) Get() {}

func (c *httpClient) Put() {}

func (c *httpClient) Post() {}

func (c *httpClient) Patch() {}

func (c *httpClient) Delete() {}
