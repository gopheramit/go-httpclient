package example

import "fmt"

type Endpoints struct {
	CurrentUserURL    string `json:"current_user_url"`
	AuthorizationsURL string `json:"authorizations_url"`
	RepositoryURL     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("http://api.github.com", nil)
	if err != nil {
		//Deal with error
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Status code: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Headers: %v", response.Headers()))
	fmt.Println(fmt.Sprintf("Body: %s", response.String()))

	var endpoints Endpoints
	err = response.UnmarshalJSON(&endpoints)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Repository Url : %s", endpoints.RepositoryURL))
	return &endpoints, nil

}
