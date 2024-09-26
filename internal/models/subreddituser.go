package models

import "time"

// SubredditUser is the user model
type SubredditUser struct {
	SubredditUserId string
	SubredditId     string
	UserId          string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Subreddit       Subreddit
	User            User
}
