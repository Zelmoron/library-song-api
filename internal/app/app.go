package app

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/database"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
	"EffectiveMobile/internal/services"
	"os"

	_ "EffectiveMobile/documentation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.app.Use(logger.New(), recover.New())
	db := database.CreateTables()
	app.postgre = postgre.New(db)
	app.services = services.New(app.postgre)
	app.endpoints = endpoints.New(app.services, app.postgre)

	app.routers()

	return app
}

func (a *App) routers() {

	a.app.Post("/song", a.endpoints.CreateSong)
	a.app.Get("/songs", a.endpoints.GetSongs)
	a.app.Get("/song-verse", a.endpoints.GetSongsWithVerses)

	a.app.Get("/info", func(c *fiber.Ctx) error {

		response := api.SongInfoResponse{
			ReleaseDate: "16.07.2006",
			Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
			Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})
	a.app.Get("/swagger/*", swagger.HandlerDefault)
}

func (a *App) Run() {
	a.app.Listen(os.Getenv("PORT"))
}
