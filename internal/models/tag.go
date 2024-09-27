package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Tag is the user model
type Tag struct {
	TagId       uuid.UUID
	SubredditId uuid.UUID
	AdminId     uuid.UUID
	Name        string
	Color       string
	IsMature    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
