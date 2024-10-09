package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// GetSubredditsTopics get the list of all subredditstopics from the database
func (m *postgresDBRepo) GetSubredditsTopics() ([]models.SubredditTopic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var subredditstopics []models.SubredditTopic
	stmt := `SELECT * FROM subreddits_topics`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return subredditstopics, err
	}
	for rows.Next() {
		var st models.SubredditTopic
		err = rows.Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
		if err != nil {
			return subredditstopics, err
		}
		subredditstopics = append(subredditstopics, st)
	}
	return subredditstopics, err
}

// InsertSubredditTopic inserts subreddits topics into the database
func (m *postgresDBRepo) InsertSubredditTopic(res models.SubredditTopic) (models.SubredditTopic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var st models.SubredditTopic
	stmt := `INSERT INTO subreddits_topics (subreddit_id, topic_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, res.SubredditId, res.TopicId).Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
	return st, err
}
