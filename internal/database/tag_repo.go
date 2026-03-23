package database

import (
	"fmt"
	"serhii/internal/model"
	"serhii/internal/utils"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TagRepo struct {
	db *sqlx.DB
}

func NewTagRepo(db *sqlx.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (tr *TagRepo) SingleByName(name string) (*model.Tag, error) {
	var tag model.Tag

	query := "SELECT * FROM tags WHERE name = ?"
	if err := tr.db.Get(&tag, query, name); err != nil {
		return nil, fmt.Errorf("tag_repo SingleByName() error: %v", err)
	}

	return &tag, nil
}

func (tr *TagRepo) ForPosts(postsIDs []int) ([]*model.Tag, error) {
	idsStr := utils.IntsToStrings(postsIDs)

	query := `
		SELECT t.id, t.name, t.color, t.title, pt.post_id
		FROM tags t 
		JOIN post_tag pt ON t.id = pt.tag_id 
		WHERE pt.post_id IN (?);
	`

	tags := make([]*model.Tag, 0, 2)
	if err := tr.db.Select(&tags, query, strings.Join(idsStr, ",")); err != nil {
		return nil, fmt.Errorf("tag_repo ForPosts() error: %v", err)
	}

	return tags, nil
}
