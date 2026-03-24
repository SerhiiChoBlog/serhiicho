package database

import (
	"fmt"
	"log"
	"serhii/internal/model"
	"serhii/internal/utils"
	"strings"

	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (pr *PostRepo) All() ([]*model.Post, error) {
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

	if err := pr.db.Select(&posts, postsQuery); err != nil {
		return nil, fmt.Errorf("post_repo All() error: %v", err)
	}

	pr.setAccessors(posts)

	return posts, nil
}

func (pr *PostRepo) AllWithTags(tagRepo *TagRepo) ([]*model.Post, error) {
	posts, err := pr.All()
	if err != nil {
		log.Fatalln(err)
	}

	tags, err := tagRepo.ForPosts(utils.ExtractIDs(posts))
	if err != nil {
		return nil, err
	}

	model.AttachTagsToPosts(tags, posts)

	return posts, nil
}

func (pr *PostRepo) Single(slug string) (*model.Post, error) {
	var post model.Post

	query := `SELECT * FROM posts WHERE slug = ?`
	if err := pr.db.Get(&post, query, slug); err != nil {
		return nil, fmt.Errorf("post_repo Single() error: %v", err)
	}

	return &post, nil
}

func (pr *PostRepo) SingleWithTagsAndSeries(
	slug string,
	tagRepo *TagRepo,
	seriesRepo *SeriesRepo,
) (*model.Post, error) {
	post, err := pr.Single(slug)
	if err != nil {
		return nil, err
	}

	tags, err := tagRepo.ForPosts([]int{post.ID})
	if err != nil {
		return nil, err
	}

	model.AttachTagsToPosts(tags, []*model.Post{post})

	series, err := seriesRepo.ForPosts([]int{post.ID})
	if err != nil {
		return nil, err
	}

	model.AttachSeriesToPosts(series, []*model.Post{post})

	return post, nil
}

func (pr *PostRepo) Latest() ([]*model.Post, error) {
	posts := make([]*model.Post, 0, 2)

	query := `
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

	if err := pr.db.Select(&posts, query); err != nil {
		return nil, fmt.Errorf("post_repo Latest() error: %v", err)
	}

	pr.setAccessors(posts)

	return posts, nil
}

func (pr *PostRepo) LatestWithTags(tagRepo *TagRepo) ([]*model.Post, error) {
	posts, err := pr.Latest()
	if err != nil {
		return nil, err
	}

	tags, err := tagRepo.ForPosts(utils.ExtractIDs(posts))
	if err != nil {
		return nil, err
	}

	model.AttachTagsToPosts(tags, posts)

	return posts, nil
}

func (pr *PostRepo) ForSeries(seriesIDs []int) ([]*model.Post, error) {
	idsStr := utils.IntsToStrings(seriesIDs)

	query := fmt.Sprintf(`
		SELECT p.*, ps.series_id
		FROM posts p
		JOIN post_series ps ON p.id = ps.post_id
		WHERE ps.series_id IN (%s)
	`, strings.Join(idsStr, ","))

	posts := make([]*model.Post, 0, 2)
	if err := pr.db.Select(&posts, query); err != nil {
		return nil, fmt.Errorf("post_repo ForSeries() error: %v", err)
	}

	return posts, nil
}

func (pr *PostRepo) ComingSoon() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)

	query := `
		SELECT id, slug, title, intro, image_sm, created_at, updated_at, is_published
		FROM posts
		WHERE is_published = false
		ORDER BY updated_at DESC
	`

	if err := pr.db.Select(&posts, query); err != nil {
		return nil, fmt.Errorf("post_repo ComingSoon() error: %v", err)
	}

	pr.setAccessors(posts)

	return posts, nil
}

func (pr *PostRepo) ComingSoonWithTagsAndSeries(
	tagRepo *TagRepo,
	seriesRepo *SeriesRepo,
) ([]*model.Post, error) {
	posts, err := pr.ComingSoon()
	if err != nil {
		return nil, err
	}

	postsIDs := utils.ExtractIDs(posts)
	tags, err := tagRepo.ForPosts(postsIDs)
	if err != nil {
		return nil, err
	}

	model.AttachTagsToPosts(tags, posts)

	series, err := seriesRepo.ForPosts(postsIDs)
	if err != nil {
		return nil, err
	}

	model.AttachSeriesToPosts(series, posts)

	return posts, nil
}

func (pr *PostRepo) setAccessors(posts []*model.Post) {
	for i := range posts {
		posts[i].SetAccessors()
	}
}
