package models

import "time"

type PostPublication struct {
	ID             int       `json:"id" db:"id"`
	PostID         int       `json:"post_id" db:"post_id"`
	PeopleNotified int       `json:"people_notified" db:"people_notified"`
	PeopleSkipped  int       `json:"people_skipped" db:"people_skipped"`
	Post           *Post     `json:"post,omitempty" db:"-"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
