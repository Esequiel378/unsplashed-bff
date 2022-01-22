package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func PhotosList(params string) (*PaginatedPhoto, error) {
	url := BASE_URL + "/photos"

	resp, err := prepareRequest(http.MethodGet, url + params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

    total, _ := strconv.Atoi(resp.Header.Get("X-Total"))
    perPage, _ := strconv.Atoi(resp.Header.Get("X-Per-Page"))

    photos := []Photo{}

	err = json.NewDecoder(resp.Body).Decode(&photos)

	paginatedPhoto := PaginatedPhoto{
        Total: total,
        TotalPages: total / perPage,
        Results: photos,
    }

	return &paginatedPhoto, err
}
