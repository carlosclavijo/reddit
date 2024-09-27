package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Subreddit is the user model
type Subreddit struct {
	SubredditId uuid.UUID
	Name        string
	Description string
	CreatedBy   uuid.UUID
	Icon        string
	Banner      string
	Privacy     string
	IsMature    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
