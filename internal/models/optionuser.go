package models

import (
	"time"
)

// OptionUser is the user model
type OptionUser struct {
	OptionUser string
	Option     string
	User       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Option     Option
	User       User
}
