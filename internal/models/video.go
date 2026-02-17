package models

import "time"

type Video struct {
	ID          string    `json:"id" db:"id"`
	PostID      *int      `json:"post_id,omitempty" db:"post_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image" db:"image"`
	URL         string    `json:"url" db:"url"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
