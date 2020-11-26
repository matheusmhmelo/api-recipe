package recipes

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/matheusmhmelo/api-recipe/internal/utils"
	"net/url"
)

func findGif(title string) (string, error) {
	endpoint := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape(title))

	apiResults, err := utils.Request(endpoint)
	if err != nil {
		return "", err
	}

	if len(apiResults) == 0 || apiResults["data"] == nil {
		err = errors.New("something wrong with the response of Giphy API")
		return "", err
	}

	giphyResult := apiResults["data"].([]interface{})
	firstGiphy := giphyResult[0].(map[string]interface{})

	if len(firstGiphy) == 0 {
		err = errors.New("error to get Giphy url")
		return "", err
	}

	return firstGiphy["url"].(string), nil
}

