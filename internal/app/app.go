package app

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/database"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
	"EffectiveMobile/internal/services"
	"os"

	"github.com/gofiber/fiber/v2"
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
	db := database.CreateTables()
	app.postgre = postgre.New(db)
	app.services = services.New(app.postgre)
	app.endpoints = endpoints.New(app.services)

	app.routers()

	return app
}

func (a *App) routers() {
	a.app.Post("/song", a.endpoints.CreateSong)
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
