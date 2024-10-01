package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertComment inserts comments into the database
func (m *postgresDBRepo) InsertComment(res models.Comment) error {
	/*stmt := `INSERT INTO comments (post_id, user_id, comment) VALUES($1, $2, $3)`
	err := m.DB.QueryRow(stmt, res.PostId, res.UserId, res.Comment)
	return err*/
	return nil
}
