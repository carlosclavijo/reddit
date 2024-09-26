package models

import (
	"time"
)

// OptionUser is the user model
type OptionUser struct {
	OptionUser string
	OptionId   string
	UserId     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Option     Option
	User       User
}
