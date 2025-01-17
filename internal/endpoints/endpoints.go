package endpoints

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/requests"
	"EffectiveMobile/internal/responses"
	"fmt"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Services interface {
	CreateSong(requests.SongRequest) (*responses.SongInfoResponse, error)
	GetSongs(*fiber.Ctx, int, int) ([]*responses.SongInfoResponse, int, int, int, int)
	GetSongsWithVerses(*fiber.Ctx, string, string, int) ([]string, []*responses.SongInfoResponse)
	UpdateSong(string, requests.UpdateRequest) error
	DeleteSong(string) error
}
type Endpoints struct {
	services Services
}

func New(services Services) *Endpoints {

	return &Endpoints{
		services: services,
	}
}

// CreateSong godoc
// @Summary Add a new song
// @Description Create a new song
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body requests.SongRequest true "Song data"
// @Success 200 {object} responses.SongInfoResponse
// @Failure 400 {object} responses.ErrorResponse400
// @Failure 404 {object} responses.ErrorResponse404
// @Failure 422 {object} responses.ErrorResponse422
// @Failure 500 {object} responses.ErrorResponse500
// @Router /song [post]
func (e *Endpoints) CreateSong(c *fiber.Ctx) error {
	// Обработчик создания песни
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
		logrus.Error("Данные не прошли валидацию")
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
// @Failure 404 {object} responses.ErrorResponse404
// @Failure 500 {object} responses.ErrorResponse500
// @Router /songs [get]
func (e *Endpoints) GetSongs(c *fiber.Ctx) error {
	//Обработчик получения записей из базы данных

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	songs, page, limit, totalCount, totalPages := e.services.GetSongs(c, page, limit)

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

// GetSongsWithVerses godoc
// @Summary Get Songs With Verses
// @Description Get song with pagination on verses
// @Tags Songs
// @Accept json
// @Produce json
// @Param song query string false "Song filter"
// @Param group query string false "Group filter"
// @Param verses query string false "Verses filter"
// @Success 200 {object} responses.SongResponse
// @Failure 400 {object} responses.ErrorResponse400
// @Failure 404 {object} responses.ErrorResponse404
// @Failure 500 {object} responses.ErrorResponse500
// @Router /song-verse [get]
func (e *Endpoints) GetSongsWithVerses(c *fiber.Ctx) error {
	//Обработчтик получения песни и куплетов в ней
	songName := c.Query("song")
	group := c.Query("group")
	versesLimit := c.QueryInt("verses", 1)

	if songName == "" || group == "" {
		logrus.Info("Песня или группа не должна/ы быть пустой")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Song/group name is required",
		})
	}

	if versesLimit < 1 {
		logrus.Info("Кол-во куплетов должно быть больше 0")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Number of verses must be greater than 0",
		})
	}

	verses, song := e.services.GetSongsWithVerses(c, songName, group, versesLimit)
	if len(verses) == 0 {
		logrus.Error("Ошибка получения песни")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to retrieve song",
		})
	}

	if versesLimit > len(verses) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("This song have only %d verses", len(verses)),
		})
	}
	response := responses.SongResponse{
		Song:   song[0].Song,
		Group:  song[0].Group,
		Verses: verses[:versesLimit],
	}

	return c.Status(http.StatusOK).JSON(response)
}

// Update godoc
// @Summary Update songs
// @Description Update songs
// @Tags Songs
// @Accept json
// @Produce json
// @Param  id   path    string  true  "song id"
// @Param song body requests.UpdateRequest true "Song data"
// @Success 200 {object} responses.UpdateResponse
// @Failure 400 {object} responses.ErrorResponse400
// @Failure 500 {object} responses.ErrorResponse500
// @Router /song/{id} [patch]
func (e *Endpoints) UpdateSong(c *fiber.Ctx) error {
	//Обработчик обновления песни
	id := c.Params("id")
	var update requests.UpdateRequest
	if err := c.BodyParser(&update); err != nil {
		errResp := responses.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request - Invalid input data",
		}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	err := e.services.UpdateSong(id, update)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to update",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Update succeeded",
	})
}

// Delete godoc
// @Summary Delete songs
// @Description Delete songs
// @Tags Songs
// @Accept json
// @Produce json
// @Param  id   path    string  true  "song id"
// @Success 200 {object} responses.DeleteResponse
// @Failure 400 {object} responses.ErrorResponse400
// @Failure 500 {object} responses.ErrorResponse500
// @Router /song/{id} [delete]
func (e *Endpoints) DeleteSong(c *fiber.Ctx) error {
	//Обработчик удаления песни
	id := c.Params("id")

	err := e.services.DeleteSong(id)
	if err != nil {
		logrus.Error(err)
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to delete",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Delete succeeded",
	})
}
