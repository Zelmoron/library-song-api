package main

import (
	"EffectiveMobile/internal/app"
	"log"

	// _ "github.com/username/reponame/docs" // замените на ваш путь к документации
	_ "EffectiveMobile/docs" // замените "myproject" на ваш модуль, указанный в go.mod

	"github.com/joho/godotenv"
)

func init() {
	//инициализация env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// @title TZ
// @version 1.0
// @description Rest API Library
// @termsOfService http://swagger.io/terms/
// @host localhost:3000
// @BasePath /
func main() {
	app := app.New()
	app.Run() //запуск сервера
}
