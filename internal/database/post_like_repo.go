package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostLikeRepo struct {
	db *sqlx.DB
}

func NewPostLikeRepo(db *sqlx.DB) *PostLikeRepo {
	return &PostLikeRepo{db: db}
}

func (plr *PostLikeRepo) CountPostLikes(postID int) (int, error) {
	var count int

	query := "SELECT COUNT(*) FROM post_likes WHERE post_id = ?"
	if err := plr.db.Get(&count, query, postID); err != nil {
		return 0, fmt.Errorf("count post_likes error in CountPostLikes(): %v", err)
	}

	return count, nil
}

func (plr *PostLikeRepo) UserLikedPost(postID, userID int) (bool, error) {
	var count int

	query := `
		SELECT COUNT(1)
		FROM post_likes
		WHERE post_id = ? AND user_id = ?
	`

	if err := plr.db.Get(&count, query, postID, userID); err != nil {
		return false, fmt.Errorf("count post_likes error in UserLikedPost(): %v", err)
	}

	return count > 0, nil
}
