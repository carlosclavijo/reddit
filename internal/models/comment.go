package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Comment is the user model
type Comment struct {
	CommentId  uuid.UUID
	PostId     uuid.UUID
	UserId     uuid.UUID
	ResponseId uuid.UUID
	Comment    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
