package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetSubredditsUsers() ([]models.SubredditUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var SubredditUsers []models.SubredditUser
	stmt := `SELECT * FROM subreddits_users`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return SubredditUsers, err
	}
	for rows.Next() {
		var su models.SubredditUser
		err = rows.Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
		if err != nil {
			return SubredditUsers, err
		}
		SubredditUsers = append(SubredditUsers, su)
	}
	return SubredditUsers, err
}

func (m *postgresDBRepo) GetSubredditUserById(subredditUserId string) (models.SubredditUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var su models.SubredditUser
	stmt := `SELECT * FROM subreddits_users WHERE subreddit_user_id = $1`
	uid, _ := uuid.FromString(subredditUserId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	return su, err
}

func (m *postgresDBRepo) GetSubredditMembers(subredditUserId string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []models.User
	stmt := `SELECT U.* FROM users U JOIN subreddits_users SU ON U.user_id = SU.user_id JOIN subreddits S ON S.subreddit_id = SU.subreddit_id WHERE SU.subreddit_id = $1`
	uid, _ := uuid.FromString(subredditUserId)
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

func (m *postgresDBRepo) GetSubredditMembersByRole(subredditUserId string, role string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []models.User
	stmt := `SELECT U.* FROM users U JOIN subreddits_users SU ON U.user_id = SU.user_id JOIN subreddits S ON S.subreddit_id = SU.subreddit_id WHERE SU.subreddit_id = $1 AND role = $2`
	uid, _ := uuid.FromString(subredditUserId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid, role)
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

func (m *postgresDBRepo) InsertSubredditUser(r models.SubredditUser) (models.SubredditUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var su models.SubredditUser
	stmt := `INSERT INTO subreddits_users(subreddit_id, user_id, role) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.SubredditId, r.UserId, r.Role).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	return su, err
}

func (m *postgresDBRepo) UpdateSubredditUser(subredditUserId string, r models.SubredditUser) (models.SubredditUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var su models.SubredditUser
	stmt := `UPDATE subreddits_users SET `
	if r.Role != "" {
		stmt += `role = $1, `
	} else {
		return su, errors.New("there is no role")
	}
	stmt += `updated_at = NOW() WHERE subreddit_user_id = $2 RETURNING *`
	uid, _ := uuid.FromString(subredditUserId)
	err := m.DB.QueryRowContext(ctx, stmt, r.Role, uid).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	return su, err
}

func (m *postgresDBRepo) DeleteSubredditUser(subredditUserId string) (models.SubredditUser, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var su models.SubredditUser
	stmt := `DELETE FROM subreddits_users  WHERE subreddit_user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(subredditUserId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	return su, err
}
