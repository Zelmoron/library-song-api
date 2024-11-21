package endpoints

import (
	"EffectiveMobile/internal/api"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Services interface {
	CreateSong(SongRequest) (*api.SongInfoResponse, error)
}
type Endpoints struct {
	services Services
}

func New(services Services) *Endpoints {
	return &Endpoints{
		services: services,
	}
}

type SongRequest struct {
	Group string `json:"group" validate:"required,min=0"`
	Song  string `json:"song" validate:"required,min=0"`
}

func (e *Endpoints) CreateSong(c *fiber.Ctx) error {
	var song SongRequest
	if err := c.BodyParser(&song); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	validate := validator.New()
	err := validate.Struct(song)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	response, err := e.services.CreateSong(song)
	if err != nil {
		var statusCode int
		var message string
		switch err {
		case api.ErrBadRequest:
			statusCode = http.StatusBadRequest
			message = err.Error()
		case api.ErrNoResponce:
			statusCode = http.StatusNotFound
			message = err.Error()
		default:
			statusCode = fiber.StatusInternalServerError
			message = "unexpected error"
		}
		return c.Status(statusCode).JSON(fiber.Map{
			"error": message,
		})
	}
	if response == nil {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.Status(http.StatusOK).JSON(response)
}
