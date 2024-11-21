package services

import (
	"EffectiveMobile/internal/api"
	"EffectiveMobile/internal/endpoints"
	"EffectiveMobile/internal/postgre"
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
