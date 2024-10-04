package dbrepo

import (
	"errors"
	"strconv"

	"github.com/carlosclavijo/reddit/internal/models"
)

// GetTopics get the list of all topics from the database
func (m *postgresDBRepo) GetTopics() ([]models.Topic, error) {
	var topics []models.Topic
	stmt := `SELECT * FROM topics`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return topics, err
	}
	for rows.Next() {
		var t models.Topic
		err = rows.Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return topics, err
		}
		t.User, err = m.GetUserById(t.UserId.String())
		if err != nil {
			return topics, err
		}
		topics = append(topics, t)
	}
	return topics, err
}

// GetTopicById gets the topic with their uuid
func (m *postgresDBRepo) GetTopicById(id string) (models.Topic, error) {
	var t models.Topic
	stmt := `SELECT * FROM topics WHERE topic_id = '` + id + `'`
	err := m.DB.QueryRow(stmt).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return t, err
	}
	t.User, err = m.GetUserById(t.UserId.String())
	return t, err
}

// GetSubTopics get the list of all subtopics from the database
func (m *postgresDBRepo) GetSubTopics(id string) ([]models.Topic, error) {
	var topics []models.Topic
	stmt := `SELECT * FROM topics WHERE sup_topic = '` + id + `'`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return topics, err
	}
	for rows.Next() {
		var t models.Topic
		err = rows.Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return topics, err
		}
		t.User, err = m.GetUserById(t.UserId.String())
		if err != nil {
			return topics, err
		}
		topics = append(topics, t)
	}
	return topics, err
}

// InsertTopic inserts topics into the database
func (m *postgresDBRepo) InsertTopic(r models.Topic) (models.Topic, error) {
	var t models.Topic
	user, err := m.GetUserById(r.UserId.String())
	if err != nil {
		return t, err
	}
	if !user.Admin {
		return t, errors.New("you can't add a topic because you're not an admin")
	}
	stmt := `INSERT INTO topics(user_id, name`
	if r.SupTopic.Valid {
		stmt += `, sup_topic, adult_content) VALUES($1, $2, $3, $4) RETURNING *`
		err := m.DB.QueryRow(stmt, r.UserId, r.Name, r.SupTopic, r.AdultContent).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
		return t, err
	}
	stmt += `, adult_content) VALUES($1, $2, $3) RETURNING *`
	err = m.DB.QueryRow(stmt, r.UserId, r.Name, r.AdultContent).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return t, err
	}
	t.User, err = m.GetUserById(r.UserId.String())
	return t, err
}

// UpdateTopic updates topic information
func (m *postgresDBRepo) UpdateTopic(id string, r models.Topic) (models.Topic, error) {
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
	stmt += `adult_content = ` + strconv.FormatBool(r.AdultContent) + `, updated_at = NOW() WHERE topic_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

// DeleteSubreddit deleters the subreddit from the database
func (m *postgresDBRepo) DeleteTopic(id string) (models.Topic, error) {
	var t models.Topic
	stmt := `DELETE FROM topics  WHERE topic_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return t, err
	}
	return t, err
}
