package models

import "time"

// Tag is the user model
type Tag struct {
	TagId       string
	SubredditId string
	AdminId     string
	Name        string
	Color       string
	IsMature    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Subreddit   Subreddit
	User        User
}
