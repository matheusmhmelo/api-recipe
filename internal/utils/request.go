package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Request(endpoint string) (map[string]interface{}, error) {
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
