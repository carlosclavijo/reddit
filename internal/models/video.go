package models

import (
	"time"
)

// Video is the user model
type Video struct {
	VideoId   string
	PostId    string
	Title     string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      Post
}
