package endpoints

import (
	"EffectiveMobile/internal/api"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Services interface {
	CreateSong(SongRequest) *api.SongInfoResponse
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
	Group string `json:"group"`
	Song  string `json:"song"`
}

func (e *Endpoints) CreateSong(c *fiber.Ctx) error {
	var song SongRequest
	if err := c.BodyParser(&song); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	response := e.services.CreateSong(song)

	return c.Status(http.StatusAccepted).JSON(response)

}
