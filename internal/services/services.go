package services

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
	"math"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Services struct {
	postgre *postgre.Repository
}

func New(postgre *postgre.Repository) *Services {
	return &Services{
		postgre: postgre,
	}
}

func (s *Services) CreateSong(song endpoints.SongRequest) (*api.SongInfoResponse, error) {

	songResp, err := api.GetInfo(song.Group, song.Song)
	// log.Printf("External_api info: \nDate:%s\nText:%s\nLink:%s", songResp.ReleaseDate, songResp.Text, songResp.Link)
	if err != nil {
		return nil, err
	}

	songResp.Group = song.Group
	songResp.Song = song.Song
	err = s.postgre.InsertSong(songResp)
	if err != nil {
		return nil, err
	}

	return songResp, nil

}

func (s *Services) GetSongs(repository *postgre.Repository, c *fiber.Ctx, page, limit int) ([]*api.SongInfoResponse, int, int, int, int) {

	filter := &postgre.SongFilter{
		Group:       c.Query("group"),
		Song:        c.Query("song"),
		ReleaseDate: c.Query("release_date"),
		Text:        c.Query("text"),
		Link:        c.Query("link"),
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	songs, totalCount, err := repository.GetSongs(*filter, page, limit)
	if err != nil {
		logrus.Errorf("Failed to retrieve songs: %v", err)
		return nil, 0, 0, 0, 0

	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return songs, page, limit, totalCount, totalPages

}

func (s *Services) GetSongsWithVerses(repository *postgre.Repository, c *fiber.Ctx, songName string, versesLimit int) []string {
	// Создаем минимальный фильтр только по названию песни
	filter := postgre.SongFilter{
		Song: songName,
	}

	// Получаем песню
	songs, _, err := repository.GetSongs(filter, 1, 1)

	if err != nil {
		return []string{}
	}

	if len(songs) == 0 {
		return []string{}
	}
	// Получаем куплеты
	verses := splitIntoVerses(songs[0].Text)

	// Ограничиваем количество куплетов
	if versesLimit > len(verses) {
		versesLimit = len(verses)
	}

	return verses
}

// splitIntoVerses разбивает текст песни на куплеты
func splitIntoVerses(text string) []string {
	// Проверяем, не пустой ли текст
	if text == "" {
		return []string{}
	}

	// Разбиваем текст на строки
	lines := strings.Split(text, "\n")
	verses := make([]string, 0)
	currentVerse := make([]string, 0)

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Если строка пустая и у нас есть накопленный куплет
		if trimmedLine == "" && len(currentVerse) > 0 {
			// Добавляем накопленный куплет в результат
			verses = append(verses, strings.Join(currentVerse, "\n"))
			// Очищаем текущий куплет для следующего
			currentVerse = make([]string, 0)
			continue
		}

		// Если строка не пустая, добавляем её к текущему куплету
		if trimmedLine != "" {
			currentVerse = append(currentVerse, trimmedLine)
		}
	}

	// Добавляем последний куплет, если он есть
	if len(currentVerse) > 0 {
		verses = append(verses, strings.Join(currentVerse, "\n"))
	}

	return verses
}
