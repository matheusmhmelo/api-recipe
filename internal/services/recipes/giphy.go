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

func findGif(title string) (string, error) {
	endpoint := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape(title))
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if resp.StatusCode != 200 {
		err = errors.New("something went wrong on Giphy API")
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var apiResults map[string]interface{}
	err = json.Unmarshal(body, &apiResults)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if len(apiResults) == 0 || apiResults["data"] == nil {
		err = errors.New("something wrong with the response of Giphy API")
		fmt.Println(err)
		return "", err
	}

	giphyResult := apiResults["data"].([]interface{})
	firstGiphy := giphyResult[0].(map[string]interface{})

	if len(firstGiphy) == 0 {
		err = errors.New("error to get Giphy url")
		fmt.Println(err)
		return "", err
	}

	return firstGiphy["url"].(string), nil
}

