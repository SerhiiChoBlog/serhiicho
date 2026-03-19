package mysql

import (
	"fmt"
	"serhii/internal/model"
	"serhii/internal/utils"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Series struct {
	db *sqlx.DB
}

func NewSeries(db *sqlx.DB) *Series {
	return &Series{db: db}
}

func (s *Series) List() ([]*model.Series, error) {
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
	if err := s.db.Select(&series, seriesQuery); err != nil {
		return nil, fmt.Errorf("Select series error in List(): %v", err)
	}

	if err := s.attachPosts(series); err != nil {
		return nil, err
	}

	return series, nil
}

func (t *Series) attachPosts(series []*model.Series) error {
	seriesIDs := utils.ExtractIDs(series)
	idsStr := utils.IntsToStrings(seriesIDs)

	query := fmt.Sprintf(`
		SELECT p.*, ps.series_id as pivot_series_id
		FROM posts p
		JOIN post_series ps ON p.id = ps.post_id
		WHERE ps.series_id IN (%s)
	`, strings.Join(idsStr, ","))

	posts := make([]model.Post, 0, 2)
	if err := t.db.Select(&posts, query); err != nil {
		return fmt.Errorf("Select posts error: %v", err)
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
