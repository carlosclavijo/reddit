package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertOptionUser inserts option user of polls into the database
func (m *postgresDBRepo) InsertOptionUser(res models.OptionUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO option_users
				(option_id, user_id)
				VALUES($1, $2)`
	_, err := m.DB.ExecContext(ctx, stmt, res.OptionId, res.UserId)
	return err
}
