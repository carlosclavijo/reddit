package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts reservation into the database
func (m *postgresDBRepo) InsertUser(res models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO users
				(username, email, password, created_at, updated_at)
				VALUES($1, $2, $3, $4, $5)`
	_, err := m.DB.ExecContext(ctx, stmt, res.Username, res.Email, res.Password, time.Now(), time.Now())
	return err
}
