package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertUser inserts users into the database
func (m *postgresDBRepo) InsertUser(res models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO users
				(username, email, password)
				VALUES($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, stmt, res.Username, res.Email, res.Password)
	return err
}
