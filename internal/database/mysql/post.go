package mysql

import (
	"fmt"
	"serhii/internal/model"
	"serhii/internal/utils"
	"strings"

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

	postsQuery := `
		SELECT id, slug, title, intro, image_sm, image_xs, created_at, read_time
			FROM posts
			WHERE is_published = true
			ORDER BY created_at DESC
			LIMIT 2
	`

	if err := p.db.Select(&posts, postsQuery); err != nil {
		return nil, fmt.Errorf("Select posts error: %v", err)
	}

	if err := p.attachTags(posts); err != nil {
		return nil, err
	}

	p.setAccessors(posts)

	return posts, nil
}

func (p *Post) attachTags(posts []*model.Post) error {
	postIDs := utils.ExtractIDs(posts)
	idsStr := utils.IntsToStrings(postIDs)

	tagsQuery := fmt.Sprintf(`
		SELECT t.id, t.name, t.color, t.title, pt.post_id as pivot_post_id
		FROM tags t 
		JOIN post_tag pt ON t.id = pt.tag_id 
		WHERE pt.post_id IN (%s)
	`, strings.Join(idsStr, ","))

	tags := make([]model.Tag, 0, 2)
	if err := p.db.Select(&tags, tagsQuery); err != nil {
		return fmt.Errorf("Select tags error: %v", err)
	}

	// Group tags by post ID
	tagsByPost := make(map[int][]model.Tag)
	for _, tag := range tags {
		tagsByPost[tag.PivotPostID] = append(tagsByPost[tag.PivotPostID], tag)
	}

	// Attach tags to posts
	for _, post := range posts {
		post.Tags = tagsByPost[post.ID]
	}

	return nil
}

func (p *Post) setAccessors(posts []*model.Post) {
	for i := range posts {
		posts[i].SetAccessors()
	}
}
