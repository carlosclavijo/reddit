package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubredditTopic inserts subreddits topics into the database
func (m *postgresDBRepo) InsertSubredditTopic(res models.SubredditTopic) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO subreddits_topics
				(subreddit_id, topic_id)
				VALUES($1, $2)`
	_, err := m.DB.ExecContext(ctx, stmt, res.SubredditId, res.TopicId)
	return err
}
