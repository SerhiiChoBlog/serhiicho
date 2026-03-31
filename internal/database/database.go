package database

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	PostRepo     *PostRepo
	PostLikeRepo *PostLikeRepo
	TagRepo      *TagRepo
	SeriesRepo   *SeriesRepo
	UserRepo     *UserRepo
	VideoRepo    *VideoRepo
	ProjectRepo  *ProjectRepo
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		PostRepo:     NewPostRepo(db),
		PostLikeRepo: NewPostLikeRepo(db),
		TagRepo:      NewTagRepo(db),
		SeriesRepo:   NewSeriesRepo(db),
		VideoRepo:    NewVideoRepo(db),
		ProjectRepo:  NewProjectRepo(db),
	}
}
