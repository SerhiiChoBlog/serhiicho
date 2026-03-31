package model

type Badge struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Icon        string `json:"icon" db:"icon"`
	Users       []User `json:"users,omitempty" db:"-"`
}

func (b *Badge) Ident() int {
	return b.ID
}
