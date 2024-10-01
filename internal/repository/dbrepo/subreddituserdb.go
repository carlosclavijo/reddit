package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubredditUser inserts subreddit users into the database
func (m *postgresDBRepo) InsertSubredditUser(r models.SubredditUser) (models.SubredditUser, error) {
	var su models.SubredditUser
	stmt := `INSERT INTO subreddits_users(subreddit_id, user_id, role) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRow(stmt, r.SubredditId, r.UserId, r.Role).Scan(&su.SubredditUserId, &su.SubredditId, &su.UserId, &su.Role, &su.CreatedAt, &su.UpdatedAt)
	/*if err != nil {
		return su, err
	}
	stmt = stmt = `SELECT user_id, username, email, post_karma, comment_karma, account_available, profile_pic, created_at, updated_at FROM users WHERE user_id = '` + r.CreatedBy.String() + `'`
	*/
	return su, err
}
