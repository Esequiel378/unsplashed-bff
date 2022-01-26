package api

import (
	"net/http"
	"os"
)

var BASE_URL = "https://api.unsplash.com"

func makeRequest(method, url string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

    accessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
    authorization := "Client-ID " + accessKey

	req.Header.Add("Authorization", authorization)

	resp, err := client.Do(req)

	return resp, err
}
