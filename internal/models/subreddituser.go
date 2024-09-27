package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// SubredditUser is the user model
type SubredditUser struct {
	SubredditUserId uuid.UUID
	SubredditId     uuid.UUID
	UserId          uuid.UUID
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
