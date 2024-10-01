package models

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

// Subreddit is the user model
type Subreddit struct {
	SubredditId uuid.UUID
	Name        string
	Description string
	CreatedBy   uuid.UUID
	Icon        sql.NullString
	Banner      sql.NullString
	Privacy     string
	IsMature    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}
