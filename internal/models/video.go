package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Video is the user model
type Video struct {
	VideoId   uuid.UUID
	PostId    uuid.UUID
	Title     string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
