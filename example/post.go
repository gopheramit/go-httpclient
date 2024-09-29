package example

import (
	"errors"
	"fmt"
	"net/http"
)

type GithubError struct {
	StatusCode  int    `json:"-"`
	Message     string `json:"message"`
	DocumentURL string `json:"documentation_url"`
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Privete     bool   `json:"private"`
}

func CreateRepository(request Repository) (*Repository, error) {
	response, err := httpClient.Post("https://api.github.com/user/repos", request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusCreated {
		var githubError GithubError
		if err := response.UnmarshalJSON(&githubError); err != nil {
			return nil, errors.New("Error unmarshalling github error response")
		}
		return nil, fmt.Errorf("Error creating repository: %s. Documentation: %s", githubError.Message, githubError.DocumentURL)
	}

	var repository Repository
	if err := response.UnmarshalJSON(&repository); err != nil {
		return nil, errors.New("Error unmarshalling repository response")
	}

	return &repository, nil

}
