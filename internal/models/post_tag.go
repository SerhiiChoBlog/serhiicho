package models

type PostTag struct {
	PostID int  `json:"post_id" db:"post_id"`
	TagID  int  `json:"tag_id" db:"tag_id"`
	IsMain bool `json:"is_main" db:"is_main"` // assuming bool, adjust if int
}
