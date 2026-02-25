package model

type Tag struct {
	ID            int            `json:"id" db:"id"`
	Title         string         `json:"title" db:"title"`
	Description   string         `json:"description" db:"description"`
	Name          string         `json:"name" db:"name"`
	Color         string         `json:"color" db:"color"`
	PostsCount    *int           `json:"posts_count,omitempty" db:"posts_count"`
	Posts         []Post         `json:"posts,omitempty" db:"-"`
	Subscriptions []Subscription `json:"subscriptions,omitempty" db:"-"`
	Pivot         *PostTag       `json:"pivot,omitempty" db:"-"`
	PivotPostID   int            `json:"pivot_post_id" db:"pivot_post_id"` // from pivot table
	PivotTagID    int            `json:"pivot_tag_id" db:"pivot_tag_id"`   // from pivot table
	PivotIsMain   int            `json:"pivot_is_main" db:"pivot_is_main"` // from pivot table (assuming int, change to bool if needed)
}

func (t *Tag) Ident() int {
	return t.ID
}
