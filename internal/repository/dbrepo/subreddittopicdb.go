package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

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

func (m *postgresDBRepo) GetSubredditsTopicById(subredditTopicId string) (models.SubredditTopic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var st models.SubredditTopic
	stmt := `SELECT * FROM subreddits_topics WHERE subreddit_topic_id = $1`
	uid, _ := uuid.FromString(subredditTopicId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
	return st, err
}

func (m *postgresDBRepo) GetSubredditsByTopicId(topicId string) ([]models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var subreddits []models.Subreddit
	stmt := `SELECT s.* FROM subreddits s JOIN subreddits_topics st ON s.subreddit_id = st.subreddit_id JOIN topics t ON t.topic_id = st.topic_id WHERE t.topic_id = $1`
	uid, _ := uuid.FromString(topicId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid)
	if err != nil {
		return subreddits, err
	}
	for rows.Next() {
		var s models.Subreddit
		err = rows.Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return subreddits, err
		}
		subreddits = append(subreddits, s)
	}
	return subreddits, err
}

func (m *postgresDBRepo) GetTopicsBySubredditId(subredditId string) ([]models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topics []models.Topic
	stmt := `SELECT t.* FROM topics t JOIN subreddits_topics st ON t.topic_id = st.topic_id JOIN subreddits s ON s.subreddit_id = st.subreddit_id WHERE  s.subreddit_id = $1`
	uid, _ := uuid.FromString(subredditId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid)
	if err != nil {
		return topics, err
	}
	for rows.Next() {
		var t models.Topic
		err = rows.Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return topics, err
		}
		topics = append(topics, t)
	}
	return topics, err
}

func (m *postgresDBRepo) InsertSubredditTopic(res models.SubredditTopic) (models.SubredditTopic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var st models.SubredditTopic
	stmt := `INSERT INTO subreddits_topics (subreddit_id, topic_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, res.SubredditId, res.TopicId).Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
	return st, err
}

func (m *postgresDBRepo) DeleteSubredditTopic(subredditTopicId string) (models.SubredditTopic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var st models.SubredditTopic
	stmt := `DELETE FROM subreddits_topics WHERE subreddit_topic_id = $1 RETURNING *`
	uid, _ := uuid.FromString(subredditTopicId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&st.SubredditTopicId, &st.SubredditId, &st.TopicId, &st.CreatedAt, &st.UpdatedAt)
	return st, err
}
