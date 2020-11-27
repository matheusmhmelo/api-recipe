package giphy

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/matheusmhmelo/api-recipe/internal/infrastructure/cache"
	"github.com/matheusmhmelo/api-recipe/internal/infrastructure/request"
	"log"
	"net/url"
)

type Find struct {
	request request.Request
}

func New() *Find {
	return &Find{request: request.New()}
}

func (f *Find) Find(title string, cache *cache.Redis) (string, error) {
	var gif string
	endpoint := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape(title))

	var errCache error
	var gifCache string
	if cache != nil {
		gifCache, errCache = cache.Get(title)
		gif = gifCache
	}

	if errCache != nil || cache == nil {
		apiResults, err := f.request.Do(endpoint)
		if err != nil {
			return "", err
		}

		if len(apiResults) == 0 || apiResults["data"] == nil || len(apiResults["data"].([]interface{})) == 0 {
			err = errors.New("something wrong with the response of Giphy API")
			return "", err
		}

		giphyResult := apiResults["data"].([]interface{})
		firstGiphy := giphyResult[0].(map[string]interface{})

		if len(firstGiphy) == 0 {
			err = errors.New("error to get Giphy url")
			return "", err
		}

		gif = firstGiphy["url"].(string)

		if cache != nil && errCache != nil && errCache.Error() == "key not found" {
			duration := cache.Db.GetTimeDuration()
			err = cache.Set(title, gif, duration)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return gif, nil
}
