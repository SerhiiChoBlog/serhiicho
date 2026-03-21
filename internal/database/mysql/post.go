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

func (t *Post) List() ([]*model.Post, error) {
	posts := make([]*model.Post, 0, 1)

	postsQuery := `
		SELECT p.id, p.slug, p.title, p.intro, p.image_sm, p.image_xs, p.created_at, p.read_time,
			-- Select post_views_count
			(SELECT COUNT(*)
			FROM post_views pv
			WHERE pv.post_id = p.id) AS post_views_count,
			-- Select comments_count
			(SELECT COUNT(*)
			FROM comments c
			WHERE c.post_id = p.id) AS comments_count
		FROM posts p
		WHERE is_published = true
		ORDER BY created_at DESC
		LIMIT 15
	`

	if err := t.db.Select(&posts, postsQuery); err != nil {
		return nil, fmt.Errorf("Select posts error in List(): %v", err)
	}

	if err := t.attachTagsToPosts(posts); err != nil {
		return nil, err
	}

	t.setAccessors(posts)

	return posts, nil
}

func (t *Post) Single(slug string) (*model.Post, error) {
	var post model.Post

	query := `SELECT * FROM posts WHERE slug = ?`
	if err := t.db.Get(&post, query, slug); err != nil {
		return nil, fmt.Errorf("Select post error in Single(): %v", err)
	}

	if err := t.attachTagsToPost(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (t *Post) Latest() ([]*model.Post, error) {
	posts := make([]*model.Post, 0, 2)

	postsQuery := `
		SELECT p.id, p.slug, p.title, p.intro, p.image_sm, p.image_xs, p.created_at, p.read_time,
			-- Select post_views_count
			(SELECT COUNT(*)
			FROM post_views pv
			WHERE pv.post_id = p.id) AS post_views_count,
			-- Select comments_count
			(SELECT COUNT(*)
			FROM comments c
			WHERE c.post_id = p.id) AS comments_count
		FROM posts p
		WHERE p.is_published = true
		ORDER BY p.created_at DESC
		LIMIT 2
	`

	if err := t.db.Select(&posts, postsQuery); err != nil {
		return nil, fmt.Errorf("Select posts error in Latest(): %v", err)
	}

	if err := t.attachTagsToPosts(posts); err != nil {
		return nil, err
	}

	t.setAccessors(posts)

	return posts, nil
}

func (t *Post) attachTagsToPosts(posts []*model.Post) error {
	postIDs := utils.ExtractIDs(posts)
	idsStr := utils.IntsToStrings(postIDs)

	query := `
		SELECT t.id, t.name, t.color, t.title, pt.post_id as pivot_post_id
		FROM tags t 
		JOIN post_tag pt ON t.id = pt.tag_id 
		WHERE pt.post_id IN (?);
	`

	tags := make([]model.Tag, 0, 2)
	if err := t.db.Select(&tags, query, strings.Join(idsStr, ",")); err != nil {
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

func (t *Post) attachTagsToPost(post *model.Post) error {
	query := `
		SELECT t.id, t.name, t.color, t.title, pt.post_id as pivot_post_id
		FROM tags t 
		JOIN post_tag pt ON t.id = pt.tag_id 
		WHERE pt.post_id = ?;
	`

	tags := make([]model.Tag, 0, 2)
	if err := t.db.Select(&tags, query, post.ID); err != nil {
		return fmt.Errorf("Select tags error: %v", err)
	}

	// Group tags by post ID
	tagsByPost := make(map[int][]model.Tag)
	for _, tag := range tags {
		tagsByPost[tag.PivotPostID] = append(tagsByPost[tag.PivotPostID], tag)
	}

	post.Tags = tagsByPost[post.ID]

	return nil
}

func (t *Post) setAccessors(posts []*model.Post) {
	for i := range posts {
		posts[i].SetAccessors()
	}
}
