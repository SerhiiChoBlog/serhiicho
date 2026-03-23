package database

import (
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Post     *PostRepo
	PostLike *PostLikeRepo
	Tag      *TagRepo
	Series   *SeriesRepo
	User     *UserRepo
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Post:     NewPostRepo(db),
		PostLike: NewPostLikeRepo(db),
		Tag:      NewTagRepo(db),
		Series:   NewSeriesRepo(db),
	}
}
