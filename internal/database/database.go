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
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		PostRepo:     NewPostRepo(db),
		PostLikeRepo: NewPostLikeRepo(db),
		TagRepo:      NewTagRepo(db),
		SeriesRepo:   NewSeriesRepo(db),
	}
}
