package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertTopic inserts topics into the database
func (m *postgresDBRepo) InsertTopic(res models.Topic) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO topics
				(name)
				VALUES($1)`
	_, err := m.DB.ExecContext(ctx, stmt, res.Name)
	return err
}
