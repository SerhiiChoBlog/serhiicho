package model

import "time"

type SocialShare struct {
	ID        int       `json:"id" db:"id"`
	PostID    int       `json:"post_id" db:"post_id"`
	VisitorID int       `json:"visitor_id" db:"visitor_id"`
	UserID    *int      `json:"user_id,omitempty" db:"user_id"`
	Social    string    `json:"social" db:"social"`
	Post      *Post     `json:"post,omitempty" db:"-"`
	Visitor   *Visitor  `json:"visitor,omitempty" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (s *SocialShare) Ident() int {
	return s.ID
}
