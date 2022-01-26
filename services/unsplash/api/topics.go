package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func TopicsList(params string) (*PaginatedTopic, error) {
	url := BASE_URL + "/topics"

	resp, err := makeRequest(http.MethodGet, url+params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	total, _ := strconv.Atoi(resp.Header.Get("X-Total"))
	perPage, _ := strconv.Atoi(resp.Header.Get("X-Per-Page"))

	topics := []Topic{}

	err = json.NewDecoder(resp.Body).Decode(&topics)

	paginatedTopic := PaginatedTopic{
		Total:      total,
		TotalPages: total / perPage,
		Results:    topics,
	}

	return &paginatedTopic, err
}

func TopicsPhotos(slug, params string) (*PaginatedTopicPhotos, error) {
	url := fmt.Sprintf("%s/topics/%s/photos", BASE_URL, slug)

	resp, err := makeRequest(http.MethodGet, url+params)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	total, _ := strconv.Atoi(resp.Header.Get("X-Total"))
	perPage, _ := strconv.Atoi(resp.Header.Get("X-Per-Page"))

	topics := []Photo{}

	err = json.NewDecoder(resp.Body).Decode(&topics)

	paginatedTopic := PaginatedTopicPhotos{
		Total:      total,
		TotalPages: total / perPage,
		Results:    topics,
	}

	return &paginatedTopic, err
}
