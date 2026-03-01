package database

import (
	"serhii/internal/database/mysql"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Post PostRepository
	Tag  TagRepository
}

func NewMySql(db *sqlx.DB) *Database {
	return &Database{
		Post: mysql.NewPost(db),
		Tag:  mysql.NewTag(db),
	}
}

type PostRepository interface {
	Latest() ([]*model.Post, error)
	List() ([]*model.Post, error)
}

type TagRepository interface {
	ByName(name string) (*model.Tag, error)
}
