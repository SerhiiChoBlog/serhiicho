package models

type PostContentImage struct {
	ID     int    `json:"id" db:"id"`
	PostID int    `json:"post_id" db:"post_id"`
	URI    string `json:"uri" db:"uri"`
	Path   string `json:"path" db:"path"`
	Post   *Post  `json:"post,omitempty" db:"-"`
}
