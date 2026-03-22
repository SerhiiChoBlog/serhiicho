package model

type Tag struct {
	ID            int            `json:"id" db:"id"`
	Title         string         `json:"title" db:"title"`
	Description   string         `json:"description" db:"description"`
	Name          string         `json:"name" db:"name"`
	Color         string         `json:"color" db:"color"`
	PostsCount    *int           `json:"posts_count,omitempty" db:"-"`
	Posts         []Post         `json:"posts,omitempty" db:"-"`
	Subscriptions []Subscription `json:"subscriptions,omitempty" db:"-"`
	PostID        int            `json:"post_id" db:"post_id"` // from pivot table
	IsMain        int            `json:"is_main" db:"is_main"` // from pivot table
}

func (t *Tag) Ident() int {
	return t.ID
}
