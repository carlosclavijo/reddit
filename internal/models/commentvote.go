package models

import (
	"time"
)

// CommentVote is the user model
type CommentVote struct {
	CommentVoteId string
	CommentId     string
	UserId        string
	Vote          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Comment       Comment
	User          User
}
