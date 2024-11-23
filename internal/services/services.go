package services

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/postgre"
	"EffectiveMobile/internal/requests"
	"EffectiveMobile/internal/responses"
	"EffectiveMobile/internal/utils"
	"math"
	"strconv"

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

func (s *Services) CreateSong(song requests.SongRequest) (*responses.SongInfoResponse, error) {
	//Логика создания песни
	songResp, err := api.GetInfo(song.Group, song.Song)
	// log.Printf("External_api info: \nDate:%s\nText:%s\nLink:%s", songResp.ReleaseDate, songResp.Text, songResp.Link)
	if err != nil {
		return nil, err
	}
	logrus.Info("Данные из API получены успешно")

	songResp.Group = song.Group
	songResp.Song = song.Song
	err, id := s.postgre.InsertSong(songResp)
	if err != nil {
		return nil, err
	}
	songResp.Id = id
	return songResp, nil

}

func (s *Services) GetSongs(c *fiber.Ctx, page, limit int) ([]*responses.SongInfoResponse, int, int, int, int) {
	//Логика получения записей
	filter := &postgre.SongFilter{
		Group:       c.Query("group"),
		Song:        c.Query("song"),
		ReleaseDate: c.Query("releaseDate"),
		Text:        c.Query("text"),
		Link:        c.Query("link"),
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	songs, totalCount, err := s.postgre.GetSongs(*filter, page, limit)
	if err != nil {
		logrus.Errorf("Failed to retrieve songs: %v", err)
		return nil, 0, 0, 0, 0

	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return songs, page, limit, totalCount, totalPages

}

func (s *Services) GetSongsWithVerses(c *fiber.Ctx, songName string, group string, versesLimit int) ([]string, []*responses.SongInfoResponse) {
	//Логика получения песни и куплетов
	filter := postgre.SongFilter{
		Song:  songName,
		Group: group,
	}

	songs, _, err := s.postgre.GetSongs(filter, 1, 1)

	if err != nil {
		return []string{}, nil
	}

	if len(songs) == 0 {
		return []string{}, nil
	}
	// Получаем куплеты
	verses := utils.SplitIntoVerses(songs[0].Text)

	return verses, songs
}

func (s *Services) UpdateSong(id string, update requests.UpdateRequest) error {
	//Логика обновления песни
	ID, err := strconv.Atoi(id) //Переводим в int
	if err != nil {
		return err
	}

	err = s.postgre.Update(ID, update)

	if err != nil {
		return err
	}

	return nil

}

func (s *Services) DeleteSong(id string) error {
	//Логика удаления песни
	ID, err := strconv.Atoi(id) //Переводим в int
	if err != nil {
		return err
	}

	err = s.postgre.Delete(ID)

	if err != nil {
		return err
	}

	return nil

}
