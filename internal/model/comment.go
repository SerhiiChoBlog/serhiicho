package model

import "time"

type Comment struct {
	ID              int           `json:"id" db:"id"`
	PostID          int           `json:"post_id" db:"post_id"`
	UserID          int           `json:"user_id" db:"user_id"`
	VisitorID       *int          `json:"visitor_id,omitempty" db:"visitor_id"`
	CommentID       *int          `json:"comment_id,omitempty" db:"comment_id"` // parent comment ID
	Comments        []*Comment    `json:"comments,omitempty" db:"-"`            // replies
	Likes           []CommentLike `json:"likes,omitempty" db:"-"`
	User            *User         `json:"user,omitempty" db:"-"`
	Post            *Post         `json:"post,omitempty" db:"-"`
	Visitor         *Visitor      `json:"visitor,omitempty" db:"-"`
	Parent          *Comment      `json:"parent,omitempty" db:"-"`
	IsSubscribed    bool          `json:"is_subscribed" db:"is_subscribed"`
	IsApproved      bool          `json:"is_approved" db:"is_approved"`
	Comment         string        `json:"comment" db:"comment"`
	PrettyCreatedAt string        `json:"pretty_created_at" db:"-"` // generated, not in DB
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
}

func (c *Comment) Ident() int {
	return c.ID
}
