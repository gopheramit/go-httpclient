package example

import (
	"fmt"
	"net/http"

	"github.com/gopheramit/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {

	client := gohttp.NewBuilder().DisableTimeouts(true).
		SetMaxIdleConnections(5).Build()

	//	commonHeaders := make(http.Header)
	//	commonHeaders.Set("Authorization", "Bearer ABC123")
	//	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status())
}

func createUser(user User) {
	//client := gohttp.New()

	headers := make(http.Header)
	//headers.Set("Authorization", "Bearer ABC123")

	response, err := githubHttpClient.Post("https://api.github.com", headers, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
}
