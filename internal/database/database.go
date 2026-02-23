package database

import (
	"serhii/internal/database/mysql"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Post PostRepository
}

func NewMySql(db *sqlx.DB) *Database {
	return &Database{
		Post: mysql.NewPost(db),
	}
}

type PostRepository interface {
	Latest() ([]*model.Post, error)
}
