package models

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
)

// User is the user model
type User struct {
	UserId           uuid.UUID
	Username         string
	Email            string
	Password         string
	PostKarma        int
	CommentKarma     int
	AccountAvailable bool
	ProfilePic       sql.NullString
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
