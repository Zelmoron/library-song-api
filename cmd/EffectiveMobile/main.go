package main

import (
	"EffectiveMobile/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func main() {
	app := app.New()
	app.Run()
}
