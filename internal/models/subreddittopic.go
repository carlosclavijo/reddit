package models

import "time"

// SubredditTopic is the user model
type SubredditTopic struct {
	SubredditTopicId string
	SubredditId      string
	TopicId          string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Subreddit        Subreddit
	Topic            Topic
}
