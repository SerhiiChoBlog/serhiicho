package model

import "time"

type User struct {
	ID              int           `json:"id" db:"id"`
	Slug            string        `json:"slug" db:"slug"`
	Name            string        `json:"name" db:"name"`
	Secret          string        `json:"secret" db:"secret"`
	Email           string        `json:"email" db:"email"`
	EmailVerifiedAt *time.Time    `json:"email_verified_at,omitempty" db:"email_verified_at"`
	Password        string        `json:"-" db:"password"` // never expose in JSON
	CommentsCount   *int          `json:"comments_count,omitempty" db:"comments_count"`
	PostLikesCount  *int          `json:"post_likes_count,omitempty" db:"post_likes_count"`
	SignInMethod    string        `json:"sign_in_method" db:"sign_in_method"` // assuming string, adjust if SignInMethod is custom type
	Avatar          string        `json:"avatar" db:"avatar"`
	RememberToken   string        `json:"-" db:"remember_token"` // never expose
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" db:"updated_at"`
	Ban             *Ban          `json:"ban,omitempty" db:"-"`
	Roles           []Role        `json:"roles,omitempty" db:"-"`
	Badges          []Badge       `json:"badges,omitempty" db:"-"`
	PostLikes       []PostLike    `json:"post_likes,omitempty" db:"-"`
	CommentLikes    []CommentLike `json:"comment_likes,omitempty" db:"-"`
	LikedPosts      []Post        `json:"liked_posts,omitempty" db:"-"`
	LikedComments   []Comment     `json:"liked_comments,omitempty" db:"-"`
	Feed            []Feed        `json:"feed,omitempty" db:"-"`
}
