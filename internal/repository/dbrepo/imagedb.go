package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertImage inserts images into the database
func (m *postgresDBRepo) InsertImage(res models.Image) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO images
				(post_id, title, url)
				VALUES($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, stmt, res.PostId, res.Title, res.Url)
	return err
}
