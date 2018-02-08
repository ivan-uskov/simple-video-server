package model

import (
	"database/sql"
	"fmt"
)

type Video struct {
	Key       string
	Title     string
	Status    int
	Duration  int
	URL       string
	Thumbnail string
}

const (
	videoStatusReady = 3
)

type VideoRepository interface {
	Add(key string, title string, url string) error
	Get(key string) (*Video, error)
}

type videoRepository struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) VideoRepository {
	return &videoRepository{db}
}

func (r *videoRepository) Add(key string, title string, url string) error {
	q := `INSERT INTO video SET video_key = ?, title = ?, status = ?, url = ?`
	rows, err := r.db.Query(q, key, title, videoStatusReady, url)
	if err == nil {
		rows.Close()
	}

	return err
}

func (r *videoRepository) Get(key string) (*Video, error) {
	q := `SELECT video_key, title, status, duration, url, thumbnail_url FROM video WHERE video_key = ?`
	rows, err := r.db.Query(q, key)
	if err == nil {
		rows.Close()
	}

	var video Video

	for rows.Next() {
		err := rows.Scan(&video.Key, &video.Title, &video.Status, &video.Duration, &video.URL, &video.Thumbnail)
		if err != nil {
			return nil, err
		}

		return &video, nil
	}

	return nil, fmt.Errorf("video not exists")
}