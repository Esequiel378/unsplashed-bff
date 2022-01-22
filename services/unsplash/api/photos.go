package api

import (
	"encoding/json"
	"net/http"
)

func PhotosList(params string) ([]Photo, error) {
	url := BASE_URL + "/photos"

	resp, err := prepareRequest(http.MethodGet, url + params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	photos := []Photo{}

	err = json.NewDecoder(resp.Body).Decode(&photos)

	return photos, err
}
