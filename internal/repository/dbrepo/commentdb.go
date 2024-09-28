package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertComment inserts comments into the database
func (m *postgresDBRepo) InsertComment(res models.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO comments
				(post_id, user_id, comment)
				VALUES($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, stmt, res.PostId, res.UserId, res.Comment)
	return err
}
