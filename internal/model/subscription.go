package model

import "time"

type Subscription struct {
	ID          int        `json:"id" db:"id"`
	UserID      *int       `json:"user_id,omitempty" db:"user_id"`
	Email       string     `json:"email" db:"email"`
	Slug        string     `json:"slug" db:"slug"`
	Tags        []Tag      `json:"tags,omitempty" db:"-"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty" db:"confirmed_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
