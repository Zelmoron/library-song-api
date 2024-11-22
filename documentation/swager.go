// Package docs Swagger документация для API песен
package docs

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

// SongRequest описывает входные данные для создания песни
// @Description Структура для создания новой песни
type SongRequest struct {
	Group       string `json:"group" example:"Beatles"`
	Song        string `json:"song" example:"Yesterday"`
	ReleaseDate string `json:"release_date" example:"1965"`
	Text        string `json:"text" example:"Текст песни"`
	Link        string `json:"link" example:"https://example.com"`
}

// SongResponse описывает ответ после создания песни
// @Description Структура ответа после создания песни
type SongResponse struct {
	ID          int    `json:"id" example:"1"`
	Group       string `json:"group" example:"Beatles"`
	Song        string `json:"song" example:"Yesterday"`
	ReleaseDate string `json:"release_date" example:"1965"`
	Text        string `json:"text" example:"Текст песни"`
	Link        string `json:"link" example:"https://example.com"`
}

// SongListResponse структура для списка песен с пагинацией
// @Description Список песен с метаданными пагинации
type SongListResponse struct {
	Songs      []SongResponse `json:"songs"`
	Page       int            `json:"page" example:"1"`
	Limit      int            `json:"limit" example:"10"`
	Total      int            `json:"total" example:"100"`
	TotalPages int            `json:"total_pages" example:"10"`
}

// @title Музыкальная библиотека API
// @version 1.0
// @description API для управления музыкальной библиотекой

// @host localhost:8080
// @BasePath /

// CreateSong godoc
// @Summary Создание новой песни
// @Description Добавление новой песни в базу данных
// @Tags songs
// @Accept json
// @Produce json
// @Param song body SongRequest true "Информация о песне"
// @Success 201 {object} SongResponse
// @Router /song [post]
func CreateSong(c *fiber.Ctx) error {
	return nil
}

// GetSongs godoc
// @Summary Получение списка песен
// @Description Возвращает список песен с пагинацией и фильтрацией
// @Tags songs
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество записей на странице" default(10)
// @Param group query string false "Фильтр по группе"
// @Param song query string false "Фильтр по названию песни"
// @Param release_date query string false "Фильтр по дате релиза"
// @Success 200 {object} SongListResponse
// @Router /songs [get]
func GetSongs(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	// Логика получения песен (пример)
	songs := []SongResponse{
		{
			ID:          1,
			Group:       "Beatles",
			Song:        "Yesterday",
			ReleaseDate: "1965",
		},
	}

	totalCount := 50
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return c.JSON(SongListResponse{
		Songs:      songs,
		Page:       page,
		Limit:      limit,
		Total:      totalCount,
		TotalPages: totalPages,
	})
}
