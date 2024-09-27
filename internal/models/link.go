package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Link is the user model
type Link struct {
	LinkId    uuid.UUID
	PostId    uuid.UUID
	Link      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
