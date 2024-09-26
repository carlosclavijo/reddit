package models

import "time"

// Post is the user model
type Post struct {
	PostId      string
	SubredditId string
	UserId      string
	Title       string
	Description string
	Type        string
	Nsfw        bool
	Brand       bool
	Votes       int
	Comments    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Subreddit   Subreddit
	User        User
}
