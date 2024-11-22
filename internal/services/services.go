package services

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
	"math"

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
	// Обновленный фильтр с новыми полями

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
