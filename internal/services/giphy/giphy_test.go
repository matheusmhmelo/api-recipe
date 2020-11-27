package giphy

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

type RequestMock struct{}

func (r *RequestMock) Do(endpoint string) (map[string]interface{}, error) {
	endpoint1 := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape("find"))
	endpoint2 := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape("findEmpty"))
	endpoint3 := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape("resultEmpty"))
	endpoint4 := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape("resultDataNil"))
	endpoint5 := fmt.Sprintf(config.Config.Giphy.Url, config.Config.Giphy.ApiKey, url.QueryEscape("resultDataEmpty"))

	var mockedResponse map[string]interface{}
	mockedError := errors.New("something wrong with the response of Giphy API")
	switch endpoint {
	case endpoint1:
		mockedResponse = map[string]interface{}{
			"data": []interface{}{
				map[string]interface{}{
					"url": "gif found",
				},
			},
		}
		mockedError = nil
	case endpoint2:
		mockedResponse = map[string]interface{}{
			"data": []interface{}{
				map[string]interface{}{},
			},
		}
		mockedError = errors.New("error to get Giphy url")
	case endpoint3:
		mockedResponse = map[string]interface{}{}
	case endpoint4:
		mockedResponse = map[string]interface{}{
			"data": nil,
		}
	case endpoint5:
		mockedResponse = map[string]interface{}{
			"data": []interface{}{},
		}
	}
	return mockedResponse, mockedError
}

func TestFind_Find(t *testing.T) {
	f := Find{request: &RequestMock{}}

	tests := []struct {
		name  string
		title string
		want  string
		err   error
	}{
		{
			"[find] Gif Found",
			"find",
			"gif found",
			nil,
		},
		{
			"[search] Gif Found but without URL",
			"findEmpty",
			"",
			errors.New("error to get Giphy url"),
		},
		{
			"[search] Empty response from API",
			"resultEmpty",
			"",
			errors.New("something wrong with the response of Giphy API"),
		},
		{
			"[search] Wrong data from API",
			"resultDataNil",
			"",
			errors.New("something wrong with the response of Giphy API"),
		},
		{
			"[search] Empty data from API",
			"resultDataEmpty",
			"",
			errors.New("something wrong with the response of Giphy API"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.Find(tt.title, nil)
			assert.Equal(t, got, tt.want)

			if tt.err != nil {
				if err == nil {
					t.Errorf("expected error: %v", tt.err.Error())
				} else if !strings.Contains(err.Error(), tt.err.Error()) {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
