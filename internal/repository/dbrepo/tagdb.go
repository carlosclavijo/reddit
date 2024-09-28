package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertTag inserts tags into the database
func (m *postgresDBRepo) InsertTag(res models.Tag) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO tags
				(subreddit_id, admin_id, name, color)
				VALUES($1, $2, $3, $4)`
	_, err := m.DB.ExecContext(ctx, stmt, res.SubredditId, res.AdminId, res.Name, res.Color)
	return err
}
