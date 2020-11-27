package recipepuppy

import (
	"errors"
	"fmt"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type RequestMock struct{}

func (r *RequestMock) Do(endpoint string) (map[string]interface{}, error) {
	endpoint1 := fmt.Sprintf(config.Config.RecipePuppy.Url, "onions%2Cgarlic", "1")
	endpoint2 := fmt.Sprintf(config.Config.RecipePuppy.Url, "wrong%2Cresponse", "1")
	endpoint3 := fmt.Sprintf(config.Config.RecipePuppy.Url, "empty%2Cresponse", "1")

	var mockedResponse map[string]interface{}
	var mockedError error
	switch endpoint {
	case endpoint1:
		mockedResponse = map[string]interface{}{
			"results": []interface{}{
				map[string]string{
					"title":       "Recipe 1",
					"ingredients": "onions,garlic",
				},
			},
		}
		mockedError = nil
	case endpoint2:
		mockedResponse = map[string]interface{}{
			"error": "something went wrong",
		}
		mockedError = errors.New("something wrong with the response of RecipePuppy API")
	case endpoint3:
		mockedResponse = nil
		mockedError = errors.New("something wrong with the response of RecipePuppy API")
	}
	return mockedResponse, mockedError
}

func TestSearch_Search(t *testing.T) {
	s := Search{request: &RequestMock{}}

	tests := []struct {
		name        string
		ingredients string
		want        []interface{}
		err         error
	}{
		{
			"[search] Recipe Found",
			"onions,garlic",
			[]interface{}{
				map[string]string{
					"title":       "Recipe 1",
					"ingredients": "onions,garlic",
				},
			},
			nil,
		},
		{
			"[search] Wrong response from API",
			"wrong,response",
			nil,
			errors.New("something wrong with the response of RecipePuppy API"),
		},
		{
			"[search] Empty response from API",
			"empty,response",
			nil,
			errors.New("something wrong with the response of RecipePuppy API"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Search(tt.ingredients, "1")
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
