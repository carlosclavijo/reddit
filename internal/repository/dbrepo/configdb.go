package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertConfig inserts configurations into the database
func (m *postgresDBRepo) InsertConfig(res models.Config) (models.Config, error) {
	var c models.Config
	stmt := `INSERT INTO configs (subreddit_id, admin_config) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRow(stmt, res.SubredditId, res.AdminConfig).Scan(&c.ConfigId, &c.SubredditId, &c.AdminConfig, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}
