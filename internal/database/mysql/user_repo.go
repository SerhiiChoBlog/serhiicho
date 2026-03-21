package mysql

import (
	"fmt"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) BySecret(secret string) (*model.User, error) {
	var user model.User

	query := "SELECT * FROM users WHERE secret = ?"
	if err := ur.db.Get(&user, query, secret); err != nil {
		return nil, fmt.Errorf("select user error in BySecret(): %v", err)
	}

	return &user, nil
}
