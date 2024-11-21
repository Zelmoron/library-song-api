package app

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/database"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
	"EffectiveMobile/internal/services"
	"math"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	app       *fiber.App
	endpoints *endpoints.Endpoints
	services  *services.Services
	postgre   *postgre.Repository
}

func New() *App {
	app := &App{}

	app.app = fiber.New()
	app.app.Use(logger.New())
	db := database.CreateTables()
	app.postgre = postgre.New(db)
	app.services = services.New(app.postgre)
	app.endpoints = endpoints.New(app.services)

	app.routers()

	return app
}

func (a *App) routers() {
	a.app.Post("/song", a.endpoints.CreateSong)
	a.app.Get("/songs", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		limit := c.QueryInt("limit", 10)

		// Create filter struct
		filter := &postgre.SongFilter{
			Group:       c.Query("group"),
			Song:        c.Query("song"),
			ReleaseDate: c.Query("release_date"),
		}

		// Validate page and limit
		if page < 1 {
			page = 1
		}
		if limit < 1 || limit > 100 {
			limit = 10
		}

		// Call repository method
		songs, totalCount, err := a.postgre.GetSongs(filter, page, limit)
		if err != nil {
			logrus.Errorf("Failed to retrieve songs: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to retrieve songs",
				"details": err.Error(),
			})
		}

		// Check if no songs found
		if len(songs) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "No songs found",
				"filter":  filter,
			})
		}

		// Calculate total pages
		totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

		// Return response
		return c.JSON(fiber.Map{
			"songs":       songs,
			"page":        page,
			"limit":       limit,
			"total":       totalCount,
			"total_pages": totalPages,
		})
	})

	a.app.Get("/info", func(c *fiber.Ctx) error {

		response := api.SongInfoResponse{
			ReleaseDate: "16.07.2006",
			Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

}

func (a *App) Run() {
	a.app.Listen(os.Getenv("PORT"))
}
