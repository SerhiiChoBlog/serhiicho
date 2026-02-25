package model

import "time"

type PostView struct {
	ID        int       `json:"id" db:"id"`
	PostID    int       `json:"post_id" db:"post_id"`
	VisitorID int       `json:"visitor_id" db:"visitor_id"`
	Amount    *int      `json:"amount,omitempty" db:"amount"`
	Visitor   *Visitor  `json:"visitor,omitempty" db:"-"`
	Post      *Post     `json:"post,omitempty" db:"-"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (p *PostView) Ident() int {
	return p.ID
}
