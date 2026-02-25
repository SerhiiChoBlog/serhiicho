package model

type Role struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Users []User `json:"users,omitempty" db:"-"`
}

func (r *Role) Ident() int {
	return r.ID
}
