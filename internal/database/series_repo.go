package database

import (
	"fmt"
	"serhii/internal/model"
	"serhii/internal/utils"

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

func (sr *SeriesRepo) PostSeries(postID int) ([]*model.Series, error) {
	query := `
		SELECT
			s.id, s.title, s.slug, s.color_from, s.color_to,
			ps.post_id, ps.part
		FROM series s
		JOIN post_series ps ON s.id = ps.series_id 
		WHERE ps.post_id = ?
	`

	series := make([]*model.Series, 0, 4)
	if err := sr.db.Select(&series, query, postID); err != nil {
		return nil, fmt.Errorf("series_repo PostSeries() error: %v", err)
	}

	return series, nil
}
