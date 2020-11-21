package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// HTTPGet some json
func HTTPGet(url string, response interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode%200 > 99 {
		return errors.New("Non 200 code returned")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}

	return nil
}
