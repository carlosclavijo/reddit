package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertSubreddit inserts subreddits into the database
func (m *postgresDBRepo) InsertSubreddit(res models.Subreddit) (models.Subreddit, error) {
	stmt := `INSERT INTO subreddits(name, description, created_by) VALUES($1, $2, $3) RETURNING subreddit_id, name, description, created_by, privacy, is_mature, created_at, updated_at`
	var subreddit models.Subreddit
	err := m.DB.QueryRow(stmt, res.Name, res.Description, res.CreatedBy, res.Privacy, res.IsMature).Scan(&res.SubredditId, &res.Name, &res.Description, &res.CreatedBy, res.Privacy, res.IsMature, res.CreatedAt, res.UpdatedAt)
	return subreddit, err
}
