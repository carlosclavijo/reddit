package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

// GetTopicUsers get the list of all topicsusers from the database
func (m *postgresDBRepo) GetTopicsUsers() ([]models.TopicUser, error) {
	var topicsusers []models.TopicUser
	stmt := `SELECT * FROM topics_users`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return topicsusers, err
	}
	for rows.Next() {
		var tu models.TopicUser
		err = rows.Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
		if err != nil {
			return topicsusers, err
		}
		topicsusers = append(topicsusers, tu)
	}
	return topicsusers, err
}

// GetTopicById gets the topic with their uuid
func (m *postgresDBRepo) GetTopicUsersById(id string) (models.TopicUser, error) {
	var tu models.TopicUser
	stmt := `SELECT * FROM topics_users WHERE topic_user_id = $1`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRow(stmt, uid).Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
	if err != nil {
		return tu, err
	}
	return tu, err
}

// InsertTopic inserts topics into the database
func (m *postgresDBRepo) InsertTopicUser(r models.TopicUser) (models.TopicUser, error) {
	var tu models.TopicUser
	stmt := `INSERT INTO topics_users (topic_id, user_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRow(stmt, r.TopicId, r.UserId).Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
	return tu, err
}

// DeleteTopicUser deletes the topicusers relation from the database
func (m *postgresDBRepo) DeleteTopicUser(id string) (models.TopicUser, error) {
	var tu models.TopicUser
	stmt := `DELETE FROM topics_users  WHERE topic_user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRow(stmt, uid).Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
	if err != nil {
		return tu, err
	}
	tu.Topic, err = m.GetTopicById(tu.TopicId.String())
	if err != nil {
		return tu, err
	}
	tu.User, err = m.GetUserById(tu.UserId.String())
	return tu, err
}
