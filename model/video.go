package model

import "database/sql"

type Video struct {
	Key       string
	Title     string
	Status    int
	Duration  int
	Url       string
	Thumbnail string
}

const (
	videoStatusReady = 3
)

type VideoRepository interface {
	Add(key string, title string, url string) error
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