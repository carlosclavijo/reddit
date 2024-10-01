package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertPoll inserts polls into the database
func (m *postgresDBRepo) InsertPoll(res models.Poll) error {
	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO polls
				(post_id)
				VALUES($1)`
	_, err := m.DB.ExecContext(ctx, stmt, res.PostId)
	return err*/
	return nil
}
