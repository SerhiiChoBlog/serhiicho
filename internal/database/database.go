package database

import (
	"serhii/internal/database/mysql"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Post   PostRepository
	Tag    TagRepository
	Series SeriesRepository
}

func NewMySql(db *sqlx.DB) *Database {
	return &Database{
		Post:   mysql.NewPost(db),
		Tag:    mysql.NewTag(db),
		Series: mysql.NewSeries(db),
	}
}

type PostRepository interface {
	Latest() ([]*model.Post, error)
	List() ([]*model.Post, error)
	Single(slug string) (*model.Post, error)
}

type SeriesRepository interface {
	List() ([]*model.Series, error)
}

type TagRepository interface {
	ByName(name string) (*model.Tag, error)
}
