package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// OptionUser is the user model
type OptionUser struct {
	OptionUserId uuid.UUID
	OptionId     uuid.UUID
	UserId       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
