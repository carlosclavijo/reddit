package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertTag inserts tags into the database
func (m *postgresDBRepo) InsertTag(res models.Tag) (models.Tag, error) {
	var t models.Tag
	stmt := `INSERT INTO tags(subreddit_id, admin_id, name, color) VALUES($1, $2, $3, $4)`
	err := m.DB.QueryRow(stmt, res.SubredditId, res.AdminId, res.Name, res.Color).Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}
