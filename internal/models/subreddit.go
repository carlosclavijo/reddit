package models

import "time"

// Subreddit is the user model
type Subreddit struct {
	SubredditId string
	Name        string
	Description string
	CreatedBy   string
	Icon        string
	Banner      string
	Privacy     string
	IsMature    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}
