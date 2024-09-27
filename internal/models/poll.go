package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Poll is the user model
type Poll struct {
	PollId    uuid.UUID
	PostId    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
