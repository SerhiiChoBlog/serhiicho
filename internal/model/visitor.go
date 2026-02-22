package model

type Visitor struct {
	ID           int           `json:"id" db:"id"`
	IP           string        `json:"ip" db:"ip"`
	Country      *string       `json:"country,omitempty" db:"country"`
	City         *string       `json:"city,omitempty" db:"city"`
	Code         *string       `json:"code,omitempty" db:"code"`
	SocialShares []SocialShare `json:"social_shares,omitempty" db:"-"`
	PostViews    []PostView    `json:"post_views,omitempty" db:"-"`
	IsBlocked    bool          `json:"is_blocked" db:"is_blocked"`
	IsRobot      bool          `json:"is_robot" db:"is_robot"`
}
