package postgre

import (
	"EffectiveMobile/internal/api"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sql.DB
}

type SongFilter struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertSong(song *api.SongInfoResponse) error {

	query := `
	INSERT INTO songs ("group", song, release_date, text, link)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	args := []any{
		song.Group,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var songID int

	err := r.db.QueryRowContext(ctx, query, args...).Scan(&songID)
	if err != nil {
		logrus.Info("Ошибка получения id, песня не добавлена")
		return err
	}

	logrus.Info("Новая песня добавлена:", songID)
	return nil
}

func (r *Repository) GetSongs(filter *SongFilter, page, limit int) ([]*api.SongInfoResponse, int, error) {
	query := `
    SELECT  "group", song, release_date, text, link
    FROM songs
    WHERE 1=1`

	var conditions []string
	var args []interface{}
	argIndex := 1

	if filter.Group != "" {
		conditions = append(conditions, fmt.Sprintf(`"group" ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.Group+"%")
		argIndex++
	}

	if filter.Song != "" {
		conditions = append(conditions, fmt.Sprintf(`song ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.Song+"%")
		argIndex++
	}

	if filter.ReleaseDate != "" {
		conditions = append(conditions, fmt.Sprintf(`release_date ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.ReleaseDate+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	countQuery := `SELECT COUNT(*) FROM songs WHERE 1=1`
	if len(conditions) > 0 {
		countQuery += " AND " + strings.Join(conditions, " AND ")
	}

	var totalCount int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)

	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	query += fmt.Sprintf(` LIMIT $%d OFFSET $%d`, argIndex, argIndex+1)
	args = append(args, limit, (page-1)*limit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var songs []*api.SongInfoResponse
	for rows.Next() {
		song := &api.SongInfoResponse{}
		err := rows.Scan(
			&song.Group,
			&song.Song,
			&song.ReleaseDate,
			&song.Text,
			&song.Link,
		)
		if err != nil {
			return nil, 0, err
		}
		songs = append(songs, song)
	}

	return songs, totalCount, nil
}