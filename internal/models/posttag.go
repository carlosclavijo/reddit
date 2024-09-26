package models

import (
	"time"
)

// PostTag is the user model
type PostTag struct {
	PostTagId string
	PostId    string
	TagId     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      Post
	Tag       Tag
}
