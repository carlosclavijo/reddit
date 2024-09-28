package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertUser inserts users into the database
func (m *postgresDBRepo) InsertUser(res models.User) (models.User, error) {
	stmt := `INSERT INTO users (username, email, password) VALUES($1, $2, $3) RETURNING user_id, username, email, password, account_available, created_at, updated_at`
	var user models.User
	err := m.DB.QueryRow(stmt, res.Username, res.Email, res.Password).Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.AccountAvailable, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
