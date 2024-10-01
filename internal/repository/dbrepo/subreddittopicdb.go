package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubredditTopic inserts subreddits topics into the database
func (m *postgresDBRepo) InsertSubredditTopic(res models.SubredditTopic) (models.SubredditTopic, error) {
	var st models.SubredditTopic
	stmt := `INSERT INTO subreddits_topics (subreddit_id, topic_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRow(stmt, res.SubredditId, res.TopicId).Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
	return st, err
}
