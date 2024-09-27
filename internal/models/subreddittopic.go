package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// SubredditTopic is the user model
type SubredditTopic struct {
	SubredditTopicId uuid.UUID
	SubredditId      uuid.UUID
	TopicId          uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
