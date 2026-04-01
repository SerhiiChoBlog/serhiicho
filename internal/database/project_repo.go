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

func (pr *ProjectRepo) All() ([]*model.Project, error) {
	var projects []*model.Project

	query := `
		SELECT title, language, description, name, stars, has_homepage, url, updated_at
		FROM projects
		ORDER BY updated_at DESC
	`

	if err := pr.db.Select(&projects, query); err != nil {
		return nil, fmt.Errorf("project_repo All() error: %v", err)
	}

	return projects, nil
}

func (pr *ProjectRepo) Single(name string) (*model.Project, error) {
	var project model.Project

	query := `SELECT * FROM projects WHERE name = ?`
	if err := pr.db.Get(&project, query, name); err != nil {
		return nil, fmt.Errorf("project_repo Single() error: %v", err)
	}

	return &project, nil
}
