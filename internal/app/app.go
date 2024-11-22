package app

import (
	"EffectiveMobile/internal/database"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
	"EffectiveMobile/internal/services"
	"os"

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
	//Настройки приложения
	app := &App{}

	app.app = fiber.New()                                    //Получаем fiber
	app.app.Use(logger.New(), recover.New())                 //Добавляем логирование и обработку паники, если такая возникнет
	db := database.CreateTables()                            //Создаем таблицы и получаем бд
	app.postgre = postgre.New(db)                            //Получаем стуктуру Repository
	app.services = services.New(app.postgre)                 //Получаем стуктуру Services
	app.endpoints = endpoints.New(app.services, app.postgre) //Получаем стуктуру Endpoints

	app.routers() //Роутеры

	return app
}

func (a *App) routers() {

	a.app.Post("/song", a.endpoints.CreateSong)              //Добавление новой песни в бд
	a.app.Get("/songs", a.endpoints.GetSongs)                //получение всех песен с пагинацией и фильтрацией
	a.app.Get("/song-verse", a.endpoints.GetSongsWithVerses) //получение песни и пагинация по куплетам
	a.app.Patch("/song/:id", a.endpoints.UpdateSong)

	// a.app.Get("/info", func(c *fiber.Ctx) error {

	// 	response := responses.SongInfoResponse{
	// 		ReleaseDate: "16.07.2006",
	// 		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
	// 		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	// 	}
	// 	return c.Status(fiber.StatusOK).JSON(response)
	// })
	a.app.Get("/swagger/*", swagger.HandlerDefault)
}

func (a *App) Run() {
	a.app.Listen(os.Getenv("PORT")) //Стартуем на порту, указаном  в env
}
