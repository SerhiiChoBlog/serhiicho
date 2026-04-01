package database

import (
	"fmt"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type ProjectRepo struct {
	db *sqlx.DB
}

func NewProjectRepo(db *sqlx.DB) *ProjectRepo {
	return &ProjectRepo{db: db}
}

func (vr *ProjectRepo) All() ([]*model.Project, error) {
	var projects []*model.Project

	query := `
		SELECT title, language, description, name, stars, has_homepage, url, updated_at
		FROM projects
		ORDER BY updated_at DESC
	`

	if err := vr.db.Select(&projects, query); err != nil {
		return nil, fmt.Errorf("project_repo All() error: %v", err)
	}

	return projects, nil
}
