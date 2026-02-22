package database

import (
	"database/sql"
	"serhii/internal/database/mysql"
	"serhii/internal/model"
)

type Database struct {
	Post PostRepository
}

func NewMySql(db *sql.DB) *Database {
	return &Database{
		Post: mysql.NewPost(db),
	}
}

type PostRepository interface {
	Latest() ([]*model.Post, error)
}
