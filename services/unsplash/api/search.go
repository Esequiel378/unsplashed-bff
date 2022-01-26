package api

import (
	"encoding/json"
	"net/http"
)

func PhotosSearch(params string) (*PaginatedPhoto, error) {
	url := BASE_URL + "/search/photos"

	resp, err := makeRequest(http.MethodGet, url+params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	photos := PaginatedPhoto{}

	err = json.NewDecoder(resp.Body).Decode(&photos)

	return &photos, err
}
