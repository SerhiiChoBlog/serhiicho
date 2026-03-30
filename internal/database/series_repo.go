package database

import (
	"fmt"
	"serhii/internal/model"
	"serhii/internal/utils"
	"strings"

	"github.com/jmoiron/sqlx"
)

type SeriesRepo struct {
	db *sqlx.DB
}

func NewSeriesRepo(db *sqlx.DB) *SeriesRepo {
	return &SeriesRepo{db: db}
}

func (sr *SeriesRepo) All() ([]*model.Series, error) {
	seriesQuery := `
		SELECT 
			s.*,
			-- Select posts_count
			(SELECT COUNT(*) FROM post_series ps WHERE ps.series_id = s.id) AS posts_count
		FROM series s
		ORDER BY s.created_at DESC
		LIMIT 5
	`

	series := make([]*model.Series, 0, 1)
	if err := sr.db.Select(&series, seriesQuery); err != nil {
		return nil, fmt.Errorf("series_repo All() error: %v", err)
	}

	return series, nil
}

func (sr *SeriesRepo) AllWithPosts(postRepo *PostRepo) ([]*model.Series, error) {
	series, err := sr.All()
	if err != nil {
		return nil, err
	}

	seriesIDs := utils.ExtractIDs(series)
	posts, err := postRepo.ForSeries(seriesIDs)
	if err != nil {
		return nil, err
	}

	if err := model.AttachPostsToSeries(posts, series); err != nil {
		return nil, err
	}

	return series, nil
}

func (sr *SeriesRepo) ForPosts(postsIDs []int) ([]*model.Series, error) {
	idsStr := utils.IntsToStrings(postsIDs)

	query := `
		SELECT
			s.id, s.title, s.slug, s.color_from, s.color_to,
			ps.post_id, ps.part
		FROM series s
		JOIN post_series ps ON s.id = ps.series_id 
		WHERE ps.post_id IN (?)
	`

	series := make([]*model.Series, 0)
	if err := sr.db.Select(&series, query, strings.Join(idsStr, ",")); err != nil {
		return nil, fmt.Errorf("series_repo ForPosts() error: %v", err)
	}

	return series, nil
}

func (sr *SeriesRepo) Single(slug string) (*model.Series, error) {
	var series model.Series

	query := `SELECT * FROM series WHERE slug = ?`
	if err := sr.db.Get(&series, query, slug); err != nil {
		return nil, fmt.Errorf("series_repo Single() error: %v", err)
	}

	return &series, nil
}

func (sr *SeriesRepo) SingleWithPosts(slug string, postRepo *PostRepo) (*model.Series, error) {
	series, err := sr.Single(slug)
	if err != nil {
		return nil, err
	}

	series.Posts, err = postRepo.ForSeries([]int{series.ID})
	if err != nil {
		return nil, err
	}

	series.PostsCount = len(series.Posts)
	series.ReadTime = calculateSeriesReadTime(series.Posts)

	return series, nil
}

func calculateSeriesReadTime(posts []*model.Post) int {
	var total int
	for i := range posts {
		total += posts[i].ReadTime
	}
	return total
}
