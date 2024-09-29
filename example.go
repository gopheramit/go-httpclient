package main

// var (
// 	httpClient = getGithubClient()
// )

// func getGithubClient() gohttp.Client {
// 	client := gohttp.NewBuilder()
// 	commonHeaders := make(http.Header)
// 	commonHeaders.Set("Authorization", "Bearer abc-123")

// 	client.SetHeaders(commonHeaders)
// 	return client
// }

// func main() {
// 	getUrls()
// }

// func getUrls() {

// 	response, err := httpClient.Get("http://api.github.com", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(response.StatusCode)
// 	bytes, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(bytes))
// }
