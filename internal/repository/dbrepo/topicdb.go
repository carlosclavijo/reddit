package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertTopic inserts topics into the database
func (m *postgresDBRepo) InsertTopic(r models.Topic) (models.Topic, error) {
	var t models.Topic
	stmt := `INSERT INTO topics(user_id, name`
	if r.SupTopic.Valid {
		stmt += `, sup_topic) VALUES($1, $2, $3) RETURNING *`
		err := m.DB.QueryRow(stmt, r.UserId, r.Name, r.SupTopic).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
		return t, err
	}
	stmt += `) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRow(stmt, r.UserId, r.Name).Scan(&t.TopicId, &t.UserId, &t.Name, &t.SupTopic, &t.AdultContent, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}
