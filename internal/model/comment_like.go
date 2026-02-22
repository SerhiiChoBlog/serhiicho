package model

import "time"

type CommentLike struct {
	ID        int       `json:"id" db:"id"`
	CommentID int       `json:"comment_id" db:"comment_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Comment   *Comment  `json:"comment,omitempty" db:"-"`
	User      *User     `json:"user,omitempty" db:"-"`
}
