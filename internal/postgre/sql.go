package postgre

import (
	"EffectiveMobile/internal/api"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Repository struct {
	db *sql.DB
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
		fmt.Printf("QueryRowContext error: %v\n", err)
		return err
	}

	log.Printf("Id песни : %d", songID)
	return nil
}
