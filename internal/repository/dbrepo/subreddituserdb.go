package dbrepo

import (
	"errors"

	"github.com/carlosclavijo/reddit/internal/models"
)

// GetSubredditsUsers get the list of all subreddit_users relations
func (m *postgresDBRepo) GetSubredditsUsers() ([]models.SubredditUser, error) {
	var SubredditUsers []models.SubredditUser
	stmt := `SELECT * FROM subreddits_users`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return SubredditUsers, err
	}
	for rows.Next() {
		var su models.SubredditUser
		err = rows.Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
		if err != nil {
			return SubredditUsers, err
		}
		su.User, err = m.GetUserById(su.UserId.String())
		if err != nil {
			return SubredditUsers, err
		}
		su.Subreddit, err = m.GetSubredditById(su.SubredditId.String())
		if err != nil {
			return SubredditUsers, err
		}
		SubredditUsers = append(SubredditUsers, su)
	}
	return SubredditUsers, err
}

// GetSubredditById gets the subreddit with their uuid
func (m *postgresDBRepo) GetSubredditUserById(id string) (models.SubredditUser, error) {
	var su models.SubredditUser
	stmt := `SELECT * FROM subreddits_users WHERE subreddit_user_id = '` + id + `'`
	err := m.DB.QueryRow(stmt).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	if err != nil {
		return su, err
	}
	su.User, err = m.GetUserById(su.UserId.String())
	if err != nil {
		return su, err
	}
	su.Subreddit, err = m.GetSubredditById(su.SubredditId.String())
	return su, err
}

func (m *postgresDBRepo) GetSubredditMembers(id string) ([]models.User, error) {
	var users []models.User
	stmt := `SELECT U.* FROM users U JOIN subreddits_users SU ON U.user_id = SU.user_id JOIN subreddits S ON S.subreddit_id = SU.subreddit_id
		WHERE SU.subreddit_id = '` + id + `'`
	rows, err := m.DB.Query(stmt)
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

func (m *postgresDBRepo) GetSubredditMembersByRole(id string, role string) ([]models.User, error) {
	var users []models.User
	stmt := `SELECT U.* FROM users U JOIN subreddits_users SU ON U.user_id = SU.user_id JOIN subreddits S ON S.subreddit_id = SU.subreddit_id
		WHERE SU.subreddit_id = '` + id + `' AND role = '` + role + `'`
	rows, err := m.DB.Query(stmt)
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

// InsertSubredditUser inserts subreddit users into the database
func (m *postgresDBRepo) InsertSubredditUser(r models.SubredditUser) (models.SubredditUser, error) {
	var su models.SubredditUser
	stmt := `INSERT INTO subreddits_users(subreddit_id, user_id, role) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRow(stmt, r.SubredditId, r.UserId, r.Role).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	if err != nil {
		return su, err

	}
	su.Subreddit, err = m.GetSubredditById(r.SubredditId.String())
	if err != nil {
		return su, err
	}
	su.User, err = m.GetUserById(r.UserId.String())
	return su, err
}

// UpdateSubredditUser updates subreddituser information
func (m *postgresDBRepo) UpdateSubredditUser(id string, r models.SubredditUser) (models.SubredditUser, error) {
	var su models.SubredditUser
	stmt := `UPDATE subreddits_users SET `
	if r.Role != "" {
		stmt += `role = '` + r.Role + `', `
	} else {
		return su, errors.New("There is no role")
	}
	stmt += `updated_at = NOW() WHERE subreddit_user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	return su, err
}

// DeleteSubredditUser deletes the subreddituser relation
func (m *postgresDBRepo) DeleteSubredditUser(id string) (models.SubredditUser, error) {
	var su models.SubredditUser
	stmt := `DELETE FROM subreddits_users  WHERE subreddit_user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	if err != nil {
		return su, err
	}
	su.User, err = m.GetUserById(su.UserId.String())
	if err != nil {
		return su, err
	}
	su.Subreddit, err = m.GetSubredditById(su.SubredditId.String())
	return su, err
}
