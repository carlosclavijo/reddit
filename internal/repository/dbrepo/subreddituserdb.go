package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubredditUser inserts subreddit users into the database
func (m *postgresDBRepo) InsertSubredditUser(res models.SubredditUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO subreddits_users
				(subreddit_id, user_id, role)
				VALUES($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, stmt, res.SubredditId, res.UserId, res.Role)
	return err
}
