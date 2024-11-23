package postgre

import (
	"EffectiveMobile/internal/requests"
	"EffectiveMobile/internal/responses"
	"EffectiveMobile/internal/utils"

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

type VerseFilter struct {
	PageSize int // количество куплетов на странице
	Page     int // номер страницы куплетов
}

type SongFilter struct {
	Group       string   `json:"group"`
	Song        string   `json:"song"`
	ReleaseDate string   `json:"release_date"`
	Text        string   `json:"text"`
	Link        string   `json:"link"`
	Verses      []string `json:"verses"`
	TotalVerses int      `json:"total_verses"`
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertSong(song *responses.SongInfoResponse) error {

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

func (r *Repository) GetSongs(filter SongFilter, page, limit int) ([]*responses.SongInfoResponse, int, error) {
	query := `
        SELECT "group", song, release_date, text, link
        FROM songs
        WHERE 1=1`

	var conditions []string
	var args []interface{}
	argIndex := 1

	// Фильтр по группе
	if filter.Group != "" {
		conditions = append(conditions, fmt.Sprintf(`"group" ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.Group+"%")
		argIndex++
	}

	// Фильтр по названию песни
	if filter.Song != "" {
		conditions = append(conditions, fmt.Sprintf(`song ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.Song+"%")
		argIndex++
	}

	// Фильтр по дате релиза
	if filter.ReleaseDate != "" {
		conditions = append(conditions, fmt.Sprintf(`release_date ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.ReleaseDate+"%")
		argIndex++
	}

	// Фильтр по тексту песни
	if filter.Text != "" {
		conditions = append(conditions, fmt.Sprintf(`text ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.Text+"%")
		argIndex++
	}

	// Фильтр по ссылке
	if filter.Link != "" {
		conditions = append(conditions, fmt.Sprintf(`link ILIKE $%d`, argIndex))
		args = append(args, "%"+filter.Link+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	// Получение общего количества записей
	countQuery := `SELECT COUNT(*) FROM songs WHERE 1=1`
	if len(conditions) > 0 {
		countQuery += " AND " + strings.Join(conditions, " AND ")
	}

	var totalCount int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)

	err := r.db.QueryRow(countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total count: %w", err)
	}

	// Добавление пагинации
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, (page-1)*limit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to query songs: %w", err)
	}
	defer rows.Close()

	var songs []*responses.SongInfoResponse
	for rows.Next() {
		song := &responses.SongInfoResponse{}
		err := rows.Scan(
			&song.Group,
			&song.Song,
			&song.ReleaseDate,
			&song.Text,
			&song.Link,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan song: %w", err)
		}
		songs = append(songs, song)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating rows: %w", err)
	}

	return songs, totalCount, nil
}

func (r *Repository) GetSongsWithVerses(filter SongFilter, verseFilter VerseFilter) ([]*SongFilter, int, error) {

	songs, totalCount, err := r.GetSongs(filter, 1, 100)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get songs: %w", err)
	}

	songsWithVerses := make([]*SongFilter, 0, len(songs))

	for _, song := range songs {

		verses := utils.SplitIntoVerses(song.Text) //разбить по куплетам

		startIdx := (verseFilter.Page - 1) * verseFilter.PageSize
		endIdx := startIdx + verseFilter.PageSize
		if startIdx >= len(verses) {
			continue
		}
		if endIdx > len(verses) {
			endIdx = len(verses)
		}

		songWithVerses := &SongFilter{
			Group:       song.Group,
			Song:        song.Song,
			ReleaseDate: song.ReleaseDate,
			Verses:      verses[startIdx:endIdx],
			TotalVerses: len(verses),
			Link:        song.Link,
		}

		songsWithVerses = append(songsWithVerses, songWithVerses)
	}

	return songsWithVerses, totalCount, nil
}

func (r *Repository) Update(id int, req requests.UpdateRequest) error {
	// Проверяем существование записи
	query := `SELECT id FROM songs WHERE id = $1`
	var existingID int
	err := r.db.QueryRow(query, id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("song with id %d not found", id)
		}
		return fmt.Errorf("failed to check song existence: %w", err)
	}

	updateQuery := `
        UPDATE songs
        SET
            "group" = $1,
            song = $2,
            release_date = $3,
            text = $4,
            link = $5
        WHERE id = $6
    `

	result, err := r.db.Exec(updateQuery,
		req.Group,
		req.Song,
		req.ReleaseDate,
		req.Text,
		req.Link,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to update song: %w", err)
	}

	// Проверяем, что запись была обновлена
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song with id %d was not updated", id)
	}

	return nil
}

func (r *Repository) Delete(id int) error {
	query := `SELECT id FROM songs WHERE id = $1`
	var existingID int
	err := r.db.QueryRow(query, id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("song with id %d not found", id)
		}
		return fmt.Errorf("failed to check song existence: %w", err)
	}

	// Выполняем удаление
	deleteQuery := `DELETE FROM songs WHERE id = $1`

	result, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("failed to delete song: %w", err)
	}

	// Проверяем, что запись была удалена
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song with id %d was not deleted", id)
	}

	return nil

}
