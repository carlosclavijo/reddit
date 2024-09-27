package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Image is the user model
type Image struct {
	ImageId   uuid.UUID
	PostId    uuid.UUID
	Title     string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
