package api

import (
	"EffectiveMobile/internal/responses"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	ErrBadRequest          = errors.New("incorrect request")
	ErrNoResponce          = errors.New("no responce from API")
	ErrInternalServerError = errors.New("InternalServerError")
)

func GetInfo(group, song string) (*responses.SongInfoResponse, error) {
	//Получение данных из внешнего API
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

	var response responses.SongInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(response)
	if err != nil {
		logrus.Error("Данные не прошли валидацию")
		return nil, ErrBadRequest
	}

	return &response, nil
}
