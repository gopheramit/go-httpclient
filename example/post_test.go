package example

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gopheramit/go-httpclient/gohttp"
)

func TestCreateRepo(t *testing.T) {
	gohttp.FlushMocks()
	repository := Repository{
		Name:        "golang intro",
		Description: "A golang intro repository",
		Privete:     false,
	}

	gohttp.AddMock(gohttp.Mock{
		Method:      http.MethodPost,
		Url:         "https://api.github.com/user/repos",
		RequestBody: `{"name":"golang intro","description":"A golang intro repository","private":false}`,
		Error:       errors.New("Timeout from github"),
	})

	repo, err := CreateRepository(repository)
	if repo != nil {
		t.Error("No repo expected when we get timeout from github")
	}

	if err == nil {
		t.Error("Error is expected when we get timeout from github")
	}

	if err.Error() != "Timeout from github" {
		t.Error("invalid error message")
	}
}
