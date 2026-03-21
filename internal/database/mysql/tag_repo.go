package mysql

import (
	"fmt"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type TagRepo struct {
	db *sqlx.DB
}

func NewTagRepo(db *sqlx.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (tr *TagRepo) ByName(name string) (*model.Tag, error) {
	var tag model.Tag

	query := "SELECT * FROM tags WHERE name = ?"
	if err := tr.db.Get(&tag, query, name); err != nil {
		return nil, fmt.Errorf("get tag error in ByName(): %v", err)
	}

	return &tag, nil
}
