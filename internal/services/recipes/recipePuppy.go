package recipes

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

func search(ingredients string, page string) ([]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf(config.Config.RecipePuppy.Url, url.QueryEscape(ingredients), page))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = errors.New("something went wrong on RecipePuppy API")
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var apiResults map[string]interface{}
	err = json.Unmarshal(body, &apiResults)
	if err != nil {
		fmt.Println(err)
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

