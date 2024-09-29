package example

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gopheramit/go-httpclient/gohttp"
)

func TestMain(m *testing.M) {
	fmt.Println("aboout to start test cases for package example")
	gohttp.StartMockServer()
	os.Exit(m.Run())
}
func TestGetEndpoints(t *testing.T) {

	gohttp.StartMockServer()
	t.Run("Test Error Fetching from github", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "http://api.github.com",
			Error:  errors.New("Error fetching from github")})

		endpoints, err := GetEndpoints()

		if err == nil {
			t.Error("Error expected")
		}
		if endpoints != nil {
			t.Error("No endpoints expected")
		}
		if err.Error() != "Error fetching from github" {
			t.Error("Error message mismatch")
		}
	})

	t.Run("Test Error Unmarshalling response", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "http://api.github.com",
			ResponseBody:       `{"current_user_url": 123}`,
			ResponseStatusCode: http.StatusOK},
		)

		endpoints, err := GetEndpoints()
		if err == nil {
			t.Error("Error expected")
		}
		if endpoints != nil {
			t.Error("No endpoints expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("Error message mismatch")
		}
	})

	t.Run("Test Success", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "http://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})
		//Execution
		endpoints, err := GetEndpoints()
		if err != nil {
			t.Error("No error expected", err)
		}
		if endpoints == nil {
			t.Error("Endpoints  were  expected")
		}
		if endpoints.CurrentUserURL != "https://api.github.com/user" {
			t.Error("Invalid current user url")
		}
	})
}
