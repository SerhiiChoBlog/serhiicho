package mysql

import (
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{db: db}
}

func (p *Post) Latest() ([]*model.Post, error) {
	posts := make([]*model.Post, 0, 2)

	query := `
		SELECT id, slug, title, intro, image_sm, image_xs, created_at, read_time
		FROM posts
		ORDER BY created_at DESC
		LIMIT 2
	`

	err := p.db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
