package models

import "time"

// User is the user model
type User struct {
	UserId           string
	Username         string
	Email            string
	Password         string
	PostKarma        int
	CommentKarma     int
	AccountAvailable bool
	ProfilePic       string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
