package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Request interface {
	Do(string) (map[string]interface{}, error)
}

type ImplementRequest struct{}

func New() *ImplementRequest {
	return &ImplementRequest{}
}

func (i *ImplementRequest) Do(endpoint string) (map[string]interface{}, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = errors.New("something went wrong on request to: " + endpoint)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResults map[string]interface{}
	err = json.Unmarshal(body, &apiResults)
	if err != nil {
		return nil, err
	}

	return apiResults, nil
}
