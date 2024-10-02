package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// GetSubredditUsers get the list of all subreddit_users relations
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
