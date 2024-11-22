package endpoints

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/postgre"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Services interface {
	CreateSong(SongRequest) (*api.SongInfoResponse, error)
	GetSongs(*postgre.Repository, *fiber.Ctx, int, int) ([]*api.SongInfoResponse, int, int, int, int)
}
type Endpoints struct {
	repository *postgre.Repository
	services   Services
}

func New(services Services, db *postgre.Repository) *Endpoints {
	if db == nil {
		panic("database repository cannot be nil")
	}
	return &Endpoints{
		services:   services,
		repository: db,
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
	logrus.Info("Данные получены")

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

func (e *Endpoints) GetSongs(c *fiber.Ctx) error {

	if e.repository == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Repository not initialized",
		})
	}
	//Пагинация (выводит 10 записей)
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	songs, page, limit, totalCount, totalPages := e.services.GetSongs(e.repository, c, page, limit)

	if len(songs) == 0 {
		logrus.Error("Failed to retrieve songs")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No songs found",
		})
	}

	return c.JSON(fiber.Map{
		"songs":       songs,
		"page":        page,
		"limit":       limit,
		"total":       totalCount,
		"total_pages": totalPages,
	})
}
