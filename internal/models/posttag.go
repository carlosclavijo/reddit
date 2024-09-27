package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// PostTag is the user model
type PostTag struct {
	PostTagId uuid.UUID
	PostId    uuid.UUID
	TagId     uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
