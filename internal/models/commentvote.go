package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// CommentVote is the user model
type CommentVote struct {
	CommentVoteId uuid.UUID
	CommentId     uuid.UUID
	UserId        uuid.UUID
	Vote          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
