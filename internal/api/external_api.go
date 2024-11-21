package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type SongInfoResponse struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

var (
	ErrBadRequest = errors.New("incorrect request")
	ErrNoResponce = errors.New("no responce from API")
)

func GetInfo(group, song string) (*SongInfoResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/info?group=%s&song=%s", os.Getenv("API_URL"), group, song))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return nil, ErrBadRequest
		} else {
			return nil, ErrNoResponce
		}
	}

	var response SongInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
