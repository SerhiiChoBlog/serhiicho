package model

import "time"

type Ban struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Comment   string    `json:"comment" db:"comment"`
	User      *User     `json:"user,omitempty" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (b *Ban) Ident() int {
	return b.ID
}
