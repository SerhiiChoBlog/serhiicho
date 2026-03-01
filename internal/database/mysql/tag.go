package mysql

import (
	"fmt"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Tag {
	return &Tag{db: db}
}

func (t *Tag) ByName(name string) (*model.Tag, error) {
	var tag model.Tag

	query := `
		SELECT *
		FROM tags
		WHERE name = ?
	`

	if err := t.db.Get(&tag, query, name); err != nil {
		return nil, fmt.Errorf("Get tag error in ByName(): %v", err)
	}

	return &tag, nil
}
