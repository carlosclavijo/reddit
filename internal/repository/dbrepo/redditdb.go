package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubreddit inserts subreddits into the database
func (m *postgresDBRepo) InsertSubreddit(res models.Subreddit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO subreddits
				(name, description, created_by)
				VALUES($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, stmt, res.Name, res.Description, res.CreatedBy)
	return err
}
