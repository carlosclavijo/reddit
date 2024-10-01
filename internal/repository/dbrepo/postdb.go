package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertPost inserts posts into the database
func (m *postgresDBRepo) InsertPost(res models.Post) error {
	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO posts
				(subreddit_id, user_id, title, description)
				VALUES($1, $2, $3, $4)`
	_, err := m.DB.ExecContext(ctx, stmt, res.SubredditId, res.UserId, res.Title, res.Description)
	return err*/
	return nil
}
