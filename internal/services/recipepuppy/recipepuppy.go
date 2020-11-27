package recipepuppy

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/matheusmhmelo/api-recipe/internal/infrastructure/request"
	"net/url"
)

type Search struct {
	request request.Request
}

func New() *Search {
	return &Search{request: request.New()}
}

func (s *Search) Search(ingredients string, page string) ([]interface{}, error) {
	endpoint := fmt.Sprintf(config.Config.RecipePuppy.Url, url.QueryEscape(ingredients), page)

	apiResults, err := s.request.Do(endpoint)
	if err != nil {
		return nil, err
	}

	if len(apiResults) == 0 || apiResults["results"] == nil {
		err = errors.New("something wrong with the response of RecipePuppy API")
		fmt.Println(err)
		return nil, err
	}
	recipesResult := apiResults["results"].([]interface{})

	return recipesResult, nil
}
