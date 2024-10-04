package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// TopicUser is the topicuser model
type TopicUser struct {
	TopicUserId uuid.UUID
	TopicId     uuid.UUID
	UserId      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Topic       Topic
	User        User
}
