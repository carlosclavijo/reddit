package models

import (
	"time"
)

// Poll is the user model
type Poll struct {
	PollId    string
	PostId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      Post
}
