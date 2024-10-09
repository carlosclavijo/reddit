package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

// GetTopics get the list of all topics from the database
func (m *postgresDBRepo) GetTopics() ([]models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topics []models.Topic
	stmt := `SELECT * FROM topics`
	rows, err := m.DB.QueryContext(ctx, stmt)
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

// GetTopicById gets the topic with their uuid
func (m *postgresDBRepo) GetTopicById(topicId string) (models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Topic
	stmt := `SELECT * FROM topics WHERE topic_id = $1`
	uid, _ := uuid.FromString(topicId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return t, err
	}
	return t, err
}

// GetSubTopics get the list of all subtopics from the database
func (m *postgresDBRepo) GetSubTopics(topicId string) ([]models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topics []models.Topic
	stmt := `SELECT * FROM topics WHERE sup_topic = $1`
	uid, _ := uuid.FromString(topicId)
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

// GetParentsTopics gets the list of all topics without parent topics
func (m *postgresDBRepo) GetParentsTopics() ([]models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topics []models.Topic
	stmt := `SELECT * FROM topics WHERE sup_topic IS NULL`
	rows, err := m.DB.QueryContext(ctx, stmt)
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

// GetTopicsByCreatorId Get all Topics created by an User
func (m *postgresDBRepo) GetTopicsByCreatorId(userId string) ([]models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topics []models.Topic
	stmt := `SELECT t.* FROM topics t JOIN users u ON t.user_id = u.user_id WHERE t.user_id = $1`
	uid, _ := uuid.FromString(userId)
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

// InsertTopic inserts topics into the database
func (m *postgresDBRepo) InsertTopic(r models.Topic) (models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Topic
	stmt := `INSERT INTO topics(user_id, name`
	if r.SupTopic.Valid {
		stmt += `, sup_topic, adult_content) VALUES($1, $2, $3, $4) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.UserId, r.Name, r.SupTopic, r.AdultContent).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
		return t, err
	}
	stmt += `, adult_content) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.UserId, r.Name, r.AdultContent).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

// UpdateTopic updates topic information
func (m *postgresDBRepo) UpdateTopic(topicId string, r models.Topic) (models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Topic
	stmt := `UPDATE topics SET `
	if r.Name != "" {
		stmt += `name = '` + r.Name + `', `
	} else {
		stmt += `name = name, `
	}
	if t.SupTopic.Valid {
		stmt += `sup_topic = '` + t.SupTopic.UUID.String() + `', `
	} else {
		stmt += `sup_topic = sup_topic, `
	}
	stmt += `adult_content = ` + strconv.FormatBool(r.AdultContent) + `, updated_at = NOW() WHERE topic_id = $1 RETURNING *`
	uid, _ := uuid.FromString(topicId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

// DeleteSubreddit deleters the subreddit from the database
func (m *postgresDBRepo) DeleteTopic(topicId string) (models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Topic
	stmt := `DELETE FROM topics  WHERE topic_id = $1 RETURNING *`
	uid, _ := uuid.FromString(topicId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return t, err
	}
	return t, err
}
