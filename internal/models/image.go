package models

import (
	"time"
)

// Image is the user model
type Image struct {
	ImageId   string
	PostId    string
	Title     string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      Post
}
