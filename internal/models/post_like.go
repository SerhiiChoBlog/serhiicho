package models

import "time"

type PostLike struct {
	ID        int       `json:"id" db:"id"`
	PostID    int       `json:"post_id" db:"post_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	User      *User     `json:"user,omitempty" db:"-"`
	Post      *Post     `json:"post,omitempty" db:"-"`
}
