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

func (vr *ProjectRepo) Latest() ([]*model.Project, error) {
	var projects []*model.Project

	query := `
		SELECT *
		FROM projects
		ORDER BY created_at DESC
		LIMIT 4
	`

	if err := vr.db.Select(&projects, query); err != nil {
		return nil, fmt.Errorf("project_repo Latest() error: %v", err)
	}

	return projects, nil
}
