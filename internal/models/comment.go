package models

import (
	"time"
)

// Comment is the user model
type Comment struct {
	CommentId  string
	PostId     string
	UserId     string
	ResponseId string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Post       Post
	User       User
}
