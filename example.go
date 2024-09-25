package gohttpclient

import "github.com/gopheramit/gohttpclient/gohttp"

func basicExample() {
	client := gohttp.NewHttpClient()
	response, err := client.Get("http://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	println(response.Status)
}
