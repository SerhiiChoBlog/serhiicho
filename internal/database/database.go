package database

import (
	"serhii/internal/database/mysql"
	"serhii/internal/model"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Post     PostRepo
	PostLike PostLikeRepo
	Tag      TagRepo
	Series   SeriesRepo
	User     UserRepo
}

func NewMySql(db *sqlx.DB) *Database {
	return &Database{
		Post:     mysql.NewPostRepo(db),
		PostLike: mysql.NewPostLikeRepo(db),
		Tag:      mysql.NewTagRepo(db),
		Series:   mysql.NewSeriesRepo(db),
	}
}

type PostRepo interface {
	Latest() ([]*model.Post, error)
	All() ([]*model.Post, error)
	Single(slug string) (*model.Post, error)
	PostsForSeries(seriesIDs []int) ([]*model.Post, error)
}

type SeriesRepo interface {
	All() ([]*model.Series, error)
	PostSeries(postID int) ([]*model.Series, error)
}

type TagRepo interface {
	ByName(name string) (*model.Tag, error)
}

type UserRepo interface {
	BySecret(secret string) (*model.User, error)
}

type PostLikeRepo interface {
	CountPostLikes(postID int) (int, error)
	UserLikedPost(postID, userID int) (bool, error)
}
