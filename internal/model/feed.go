package model

import "time"

type Feed struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	BadgeID   *int      `json:"badge_id,omitempty" db:"badge_id"`
	CommentID *int      `json:"comment_id,omitempty" db:"comment_id"`
	PostID    *int      `json:"post_id,omitempty" db:"post_id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	Type      string    `json:"type" db:"type"` // assuming FeedType is string, adjust if custom type
	Date      time.Time `json:"date" db:"date"`
	User      *User     `json:"user,omitempty" db:"-"`
	Comment   *Comment  `json:"comment,omitempty" db:"-"`
	Post      *Post     `json:"post,omitempty" db:"-"`
	Badge     *Badge    `json:"badge,omitempty" db:"-"`
}
