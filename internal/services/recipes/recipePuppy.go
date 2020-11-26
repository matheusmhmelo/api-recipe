package recipes

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/matheusmhmelo/api-recipe/internal/utils"
	"net/url"
)

func search(ingredients string, page string) ([]interface{}, error) {
	endpoint := fmt.Sprintf(config.Config.RecipePuppy.Url, url.QueryEscape(ingredients), page)

	apiResults, err := utils.Request(endpoint)
	if err != nil {
		return nil, err
	}

	if len(apiResults) == 0 {
		err = errors.New("something wrong with the response of RecipePuppy API")
		fmt.Println(err)
		return nil, err
	}
	recipesResult := apiResults["results"].([]interface{})

	return recipesResult, nil
}

