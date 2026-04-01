package model

import "time"

type Project struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Title       string    `json:"title" db:"title"`
	URL         string    `json:"url" db:"url"`
	Description string    `json:"description" db:"description"`
	Content     string    `json:"content" db:"content"`
	Language    string    `json:"language" db:"language"`
	Stars       int       `json:"stars" db:"stars"`
	HasHomepage bool      `json:"has_homepage" db:"has_homepage"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (p *Project) Ident() int {
	return p.ID
}
