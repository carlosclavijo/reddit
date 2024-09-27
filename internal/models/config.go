package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Config is the user model
type Config struct {
	ConfigId       uuid.UUID
	SubredditId    uuid.UUID
	AdminConfig    uuid.UUID
	IsAvailable    bool
	IsLocked       bool
	TextAvailable  bool
	ImageAvailable bool
	VideoAvailable bool
	LinkAvailable  bool
	PollAvailable  bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
