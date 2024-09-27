package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Option is the user model
type Option struct {
	OptionId  uuid.UUID
	PollId    uuid.UUID
	Value     string
	Votes     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
