package mysql

import (
	"database/sql"
	"fmt"
	"serhii/internal/model"
	"time"
)

type Post struct {
	db *sql.DB
}

func NewPost(db *sql.DB) *Post {
	return &Post{db: db}
}

func (p *Post) Latest() ([]*model.Post, error) {
	posts := make([]*model.Post, 0, 2)

	query := `
		SELECT id, slug, title, intro, image_sm, image_xs, created_at, read_time
		FROM posts
		ORDER BY created_at
		DESC
		LIMIT 2;
	`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &model.Post{}
		var createdAt string

		err := rows.Scan(
			&post.ID,
			&post.Slug,
			&post.Title,
			&post.Intro,
			&post.ImageSm,
			&post.ImageXs,
			&createdAt,
			&post.ReadTime,
		)

		if err != nil {
			return nil, err
		}

		post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("-------> posts %#v\n", posts)
	fmt.Printf("-------> post[0] %#v\n", posts[0])
	fmt.Printf("-------> post[1] %#v\n", posts[1])

	return posts, nil
}
