package models

type Keyword struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Posts []Post `json:"posts,omitempty" db:"-"`
}
