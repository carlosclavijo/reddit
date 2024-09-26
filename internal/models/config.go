package models

import "time"

// Config is the user model
type Config struct {
	ConfigId       string
	SubredditId    string
	AdminConfig    string
	IsAvailable    bool
	IsLocked       bool
	TextAvailable  bool
	ImageAvailable bool
	VideoAvailable bool
	LinkAvailable  bool
	PollAvailable  bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Subreddit      Subreddit
	User           User
}
