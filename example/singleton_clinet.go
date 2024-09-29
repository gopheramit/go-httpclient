package example

import (
	"net/http"
	"time"

	"github.com/gopheramit/go-httpclient/gohttp"
	"github.com/gopheramit/go-httpclient/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(5 * time.Second).
		Build()

	return client
}
