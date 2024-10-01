package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// GetUser gets user by id from the database
func (m *postgresDBRepo) GetUser(id string) (models.User, error) {
	var u models.User
	stmt := `SELECT * FROM users WHERE user_id = '` + id + `'`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.CreatedAt, &u.UpdatedAt)
	u.Password = "restricted"
	return u, err

}

// InsertUser inserts users into the database
func (m *postgresDBRepo) InsertUser(r models.User) (models.User, error) {
	var u models.User
	stmt := `INSERT INTO users (username, email, password`
	if r.ProfilePic.Valid {
		stmt += `, profile_pic) VALUES($1, $2, $3, $4) RETURNING *`
		err := m.DB.QueryRow(stmt, r.Username, r.Email, r.Password, r.ProfilePic).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.CreatedAt, &u.UpdatedAt)
		return u, err
	}
	stmt += `) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRow(stmt, r.Username, r.Email, r.Password).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.CreatedAt, &u.UpdatedAt)
	u.Password = "restricted"
	return u, err
}
