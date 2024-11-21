package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

type SongInfoResponse struct {
	Group       string `json:"group" `
	Song        string `json:"song" `
	ReleaseDate string `json:"releaseDate" validate:"required,min=0"`
	Text        string `json:"text" validate:"required,min=0"`
	Link        string `json:"link" validate:"required,min=0"`
}

var (
	ErrBadRequest          = errors.New("incorrect request")
	ErrNoResponce          = errors.New("no responce from API")
	ErrInternalServerError = errors.New("InternalServerError")
)

func GetInfo(group, song string) (*SongInfoResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/info?group=%s&song=%s", os.Getenv("API_URL"), group, song))
	if err != nil {
		return nil, ErrInternalServerError
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

	validate := validator.New()
	err = validate.Struct(response)
	if err != nil {
		log.Info("Полученые API данные не прошли валидацию, скорее всего пустая строка")
		return nil, ErrBadRequest
	}
	return &response, nil
}
