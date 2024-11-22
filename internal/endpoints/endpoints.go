package endpoints

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/postgre"
	"EffectiveMobile/internal/requests"
	"EffectiveMobile/internal/responses"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Services interface {
	CreateSong(requests.SongRequest) (*responses.SongInfoResponse, error)
	GetSongs(*postgre.Repository, *fiber.Ctx, int, int) ([]*responses.SongInfoResponse, int, int, int, int)
	GetSongsWithVerses(*postgre.Repository, *fiber.Ctx, string, int) []string
	UpdateSong(*postgre.Repository, string, requests.UpdateRequest) error
	DeleteSong(*postgre.Repository, string) error
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

// CreateSong godoc
// @Summary Add a new song
// @Description Create a new song entry
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body requests.SongRequest true "Song data"
// @Success 200 {object} responses.SongInfoResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /song [post]
func (e *Endpoints) CreateSong(c *fiber.Ctx) error {
	var song requests.SongRequest
	if err := c.BodyParser(&song); err != nil {
		errResp := responses.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request - Invalid input data",
		}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	logrus.Info("Данные получены")
	validate := validator.New()
	if err := validate.Struct(song); err != nil {
		errResp := responses.ErrorResponse{
			Code:    fiber.StatusUnprocessableEntity, // 422
			Message: "Unprocessable Entity - Validation failed",
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errResp)
	}

	response, err := e.services.CreateSong(song)
	if err != nil {
		var errResp responses.ErrorResponse

		switch err {
		case api.ErrBadRequest:
			errResp = responses.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Bad Request - Invalid input data",
			}
			return c.Status(fiber.StatusBadRequest).JSON(errResp)

		case api.ErrNoResponce:
			errResp = responses.ErrorResponse{
				Code:    fiber.StatusNotFound,
				Message: "Not Found - Song not found",
			}
			return c.Status(fiber.StatusNotFound).JSON(errResp)

		default:
			errResp = responses.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: "Internal Server Error",
			}
			return c.Status(fiber.StatusInternalServerError).JSON(errResp)
		}
	}

	if response == nil {
		errResp := responses.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Not Found - Song not found",
		}
		return c.Status(http.StatusNotFound).JSON(errResp)
	}

	return c.Status(http.StatusOK).JSON(response)
}

// http://localhost:3000/songs?page=1&limit=4&group=f

// GetSongs godoc
// @Summary Get songs
// @Description Get songs with filtr and pagination
// @Tags Songs
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Param group query string false "Group filter"
// @Param song query string false "Song filter"
// @Param releaseDate query string false "releaseDate filter"
// @Param text query string false "Text filter"
// @Param link query string false "Link filter"
// @Success 200 {object} responses.SongsPaginationResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /songs [get]
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

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"songs":       songs,
		"page":        page,
		"limit":       limit,
		"total":       totalCount,
		"total_pages": totalPages,
	})
}

// http://localhost:3000/song-verse?song=f&verses=5
func (e *Endpoints) GetSongsWithVerses(c *fiber.Ctx) error {
	if e.repository == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Repository not initialized",
		})
	}

	// Получаем только название песни и количество куплетов из URL
	songName := c.Query("song")
	versesLimit := c.QueryInt("verses", 5) // По умолчанию 5 куплетов

	// Проверяем обязательные параметры
	if songName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Song name is required",
		})
	}

	if versesLimit < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Number of verses must be greater than 0",
		})
	}

	verses := e.services.GetSongsWithVerses(e.repository, c, songName, versesLimit)
	if len(verses) == 0 {
		logrus.Error("Failed to retrieve song")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to retrieve song",
		})
	}
	// Формируем упрощенный ответ
	response := responses.SongResponse{
		Song:   songName,
		Verses: verses[:versesLimit],
	}

	return c.JSON(response)
}

func (e *Endpoints) UpdateSong(c *fiber.Ctx) error {

	id := c.Params("id")
	var update requests.UpdateRequest
	if err := c.BodyParser(&update); err != nil {
		errResp := responses.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request - Invalid input data",
		}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	err := e.services.UpdateSong(e.repository, id, update)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusNotModified).JSON(responses.ErrorResponse{
			Code:    http.StatusNotModified,
			Message: "Failed to update",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Update succeeded",
	})
}

func (e *Endpoints) DeleteSong(c *fiber.Ctx) error {

	id := c.Params("id")

	err := e.services.DeleteSong(e.repository, id)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Delete succeeded",
	})
}
