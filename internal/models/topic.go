package models

import "time"

// Topic is the user model
type Topic struct {
	TopicId      string
	UserId       string
	Name         string
	SupTopic     int
	AdultContent bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User
}
