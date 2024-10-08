package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertConfig inserts configurations into the database
func (m *postgresDBRepo) InsertConfig(res models.Config) (models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var c models.Config
	stmt := `INSERT INTO configs (subreddit_id, admin_config) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, res.SubredditId, res.AdminConfig).Scan(&c.ConfigId, &c.SubredditId, &c.AdminConfig, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}
