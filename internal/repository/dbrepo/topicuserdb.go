package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetTopicsUsers() ([]models.TopicUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topicsusers []models.TopicUser
	stmt := `SELECT * FROM topics_users`
	rows, err := m.DB.QueryContext(ctx, stmt)
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

func (m *postgresDBRepo) GetTopicUsersById(topicUserId string) (models.TopicUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var tu models.TopicUser
	stmt := `SELECT * FROM topics_users WHERE topic_user_id = $1`
	uid, _ := uuid.FromString(topicUserId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
	return tu, err
}

func (m *postgresDBRepo) GetTopicsByUserId(userId string) ([]models.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var topics []models.Topic
	stmt := `SELECT T.* FROM topics T JOIN topics_users TU ON T.topic_id = TU.topic_id JOIN users U ON U.user_id = TU.user_id WHERE TU.user_id = $1`
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

func (m *postgresDBRepo) GetUsersByTopicId(topicId string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []models.User
	stmt := `SELECT U.* FROM users U JOIN topics_users TU ON U.user_id = TU.user_id JOIN topics T ON T.topic_id = TU.topic_id WHERE TU.topic_id = $1`
	uid, _ := uuid.FromString(topicId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, err
}

func (m *postgresDBRepo) InsertTopicUser(r models.TopicUser) (models.TopicUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var tu models.TopicUser
	stmt := `INSERT INTO topics_users (topic_id, user_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.TopicId, r.UserId).Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
	return tu, err
}

func (m *postgresDBRepo) DeleteTopicUser(topicUserId string) (models.TopicUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var tu models.TopicUser
	stmt := `DELETE FROM topics_users  WHERE topic_user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(topicUserId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&tu.TopicUserId, &tu.TopicId, &tu.UserId, &tu.CreatedAt, &tu.UpdatedAt)
	return tu, err
}
