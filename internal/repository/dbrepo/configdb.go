package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertConfig inserts configurations into the database
func (m *postgresDBRepo) InsertConfig(res models.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO configs
				(subreddit_id, admin_config)
				VALUES($1, $2)`
	_, err := m.DB.ExecContext(ctx, stmt, res.SubredditId, res.AdminConfig)
	return err
}
