package models

import (
	"time"
)

// Option is the user model
type Option struct {
	OptionId  string
	PollId    string
	Value     string
	Votes     int
	CreatedAt time.Time
	UpdatedAt time.Time
	Poll      Poll
}
