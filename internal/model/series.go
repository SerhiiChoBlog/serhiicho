package model

import "time"

type Series struct {
	ID          int       `json:"id" db:"id"`
	Slug        string    `json:"slug" db:"slug"`
	Title       string    `json:"title" db:"title"`
	Intro       string    `json:"intro" db:"intro"`
	Description string    `json:"description" db:"description"`
	ImageXs     string    `json:"image_xs" db:"image_xs"`
	ImageSm     string    `json:"image_sm" db:"image_sm"`
	ImageLg     string    `json:"image_lg" db:"image_lg"`
	ColorFrom   string    `json:"color_from" db:"color_from"`
	PostsCount  *int      `json:"posts_count,omitempty" db:"posts_count"` // NULLable
	ColorTo     string    `json:"color_to" db:"color_to"`
	ReadTime    int       `json:"read_time" db:"read_time"`
	Posts       []Post    `json:"posts" db:"-"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
