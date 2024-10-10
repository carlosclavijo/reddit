package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Post is the user model
type Post struct {
	PostId      uuid.UUID
	SubredditId uuid.UUID
	UserId      uuid.UUID
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
