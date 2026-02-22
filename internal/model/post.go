package model

import "time"

type Post struct {
	ID                  int                    `json:"id" db:"id"`
	Title               string                 `json:"title" db:"title"`
	Slug                string                 `json:"slug" db:"slug"`
	Intro               *string                `json:"intro,omitempty" db:"intro"`
	Content             *string                `json:"content,omitempty" db:"content"`
	IsPublished         bool                   `json:"is_published" db:"is_published"`
	ImageXs             string                 `json:"image_xs" db:"image_xs"`
	ImageSm             string                 `json:"image_sm" db:"image_sm"`
	ImageLg             string                 `json:"image_lg" db:"image_lg"`
	PostViewsCount      *int                   `json:"post_views_count,omitempty" db:"post_views_count"`
	SocialSharesCount   *int                   `json:"social_shares_count,omitempty" db:"social_shares_count"`
	CommentsSharesCount *int                   `json:"comments_shares_count,omitempty" db:"comments_shares_count"`
	PostLikesCount      *int                   `json:"post_likes_count,omitempty" db:"post_likes_count"`
	ReadTime            int                    `json:"read_time" db:"read_time"`
	Social              *[]map[string][]string `json:"social,omitempty" db:"social"`
	CreatedAt           time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at" db:"updated_at"`
	PublishedAt         time.Time              `json:"published_at" db:"published_at"`
	Series              []Series               `json:"series" db:"-"`
	PostLikes           []PostLike             `json:"post_likes" db:"-"`
	Video               *Video                 `json:"video,omitempty" db:"-"`
	Images              []PostContentImage     `json:"images" db:"-"`
	SocialShares        []SocialShare          `json:"social_shares" db:"-"`
	PostViews           []PostView             `json:"post_views" db:"-"`
	Tags                []Tag                  `json:"tags" db:"-"`
	Keywords            []Keyword              `json:"keywords" db:"-"`
	PostPublication     *PostPublication       `json:"post_publication,omitempty" db:"-"`
	Comments            []Comment              `json:"comments" db:"-"`
}
