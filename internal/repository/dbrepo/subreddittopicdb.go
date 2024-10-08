package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubredditTopic inserts subreddits topics into the database
func (m *postgresDBRepo) InsertSubredditTopic(res models.SubredditTopic) (models.SubredditTopic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var st models.SubredditTopic
	stmt := `INSERT INTO subreddits_topics (subreddit_id, topic_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, res.SubredditId, res.TopicId).Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
	return st, err
}
