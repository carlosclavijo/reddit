package models

import (
	"time"
)

// Link is the user model
type Link struct {
	LinkId    string
	PostId    string
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      Post
}
