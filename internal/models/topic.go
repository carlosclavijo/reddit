package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Topic is the user model
type Topic struct {
	TopicId      uuid.UUID
	UserId       uuid.UUID
	Name         string
	SupTopic     uuid.NullUUID
	AdultContent bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User
}
