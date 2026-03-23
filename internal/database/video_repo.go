package database

import (
	"fmt"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type VideoRepo struct {
	db *sqlx.DB
}

func NewVideo(db *sqlx.DB) *VideoRepo {
	return &VideoRepo{db: db}
}

func (vr *VideoRepo) Latest() ([]*model.Video, error) {
	var videos []*model.Video

	query := `
		SELECT *
		FROM videos
		ORDER BY created_at DESC
		LIMIT 4
	`

	if err := vr.db.Select(&videos, query); err != nil {
		return nil, fmt.Errorf("video_repo Latest() error: %v", err)
	}

	return videos, nil
}
