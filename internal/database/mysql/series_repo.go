package mysql

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

func (sr *SeriesRepo) List() ([]*model.Series, error) {
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
		return nil, fmt.Errorf("select series error in List(): %v", err)
	}

	if err := sr.attachPosts(series); err != nil {
		return nil, err
	}

	return series, nil
}

func (sr *SeriesRepo) attachPosts(series []*model.Series) error {
	seriesIDs := utils.ExtractIDs(series)
	idsStr := utils.IntsToStrings(seriesIDs)

	query := fmt.Sprintf(`
		SELECT p.*, ps.series_id as pivot_series_id
		FROM posts p
		JOIN post_series ps ON p.id = ps.post_id
		WHERE ps.series_id IN (%s)
	`, strings.Join(idsStr, ","))

	posts := make([]model.Post, 0, 2)
	if err := sr.db.Select(&posts, query); err != nil {
		return fmt.Errorf("select posts error in attachPosts: %v", err)
	}

	// Appends posts to series
	for _, s := range series {
		for _, post := range posts {
			if post.PivotSeriesID != s.ID {
				continue
			}
			s.Posts = append(s.Posts, post)
		}
	}

	return nil
}
