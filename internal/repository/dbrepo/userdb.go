package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// GetUsers get the list of all users from the database
func (m *postgresDBRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	stmt := `SELECT * FROM users`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, err
}

// GetUser gets user by id from the database
func (m *postgresDBRepo) GetUserById(id string) (models.User, error) {
	var u models.User
	stmt := `SELECT * FROM users WHERE user_id = '` + id + `'`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	u.Password = "restricted"
	return u, err

}

// InsertUser inserts users into the database
func (m *postgresDBRepo) InsertUser(r models.User) (models.User, error) {
	var u models.User
	stmt := `INSERT INTO users (username, email, password`
	if r.ProfilePic.Valid {
		stmt += `, profile_pic) VALUES($1, $2, $3, $4) RETURNING *`
		err := m.DB.QueryRow(stmt, r.Username, r.Email, r.Password, r.ProfilePic).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		return u, err
	}
	stmt += `) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRow(stmt, r.Username, r.Email, r.Password).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	u.Password = "restricted"
	return u, err
}

// UpdateUser updates user information
func (m *postgresDBRepo) UpdateUser(id string, r models.User) (models.User, error) {
	var u models.User
	stmt := `UPDATE users SET `
	if r.Username != "" {
		stmt += `username = '` + r.Username + `', `
	} else {
		stmt += `username = username, `
	}
	if r.Email != "" {
		stmt += `email = '` + r.Email + `', `
	} else {
		stmt += `email = email, `
	}
	if r.Password != "" {
		stmt += `password = '` + r.Password + `', `
	} else {
		stmt += `password = password`
	}
	if r.ProfilePic.String != "" {
		stmt += `profile_pic = '` + r.ProfilePic.String + `' `
	} else if !r.ProfilePic.Valid {
		stmt += `profile_pic = NULL `
	} else {
		stmt += `profile_pic = profile_pic `
	}
	stmt += `, updated_at = NOW() WHERE user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AddUserPostKarma(id string) (models.User, error) {
	var u models.User
	stmt := `UPDATE users SET post_karma = post_karma + 1, updated_at = NOW() WHERE user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AddUserCommentKarma(id string) (models.User, error) {
	var u models.User
	stmt := `UPDATE users SET comment_karma = comment_karma + 1, updated_at = NOW() WHERE user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AdminUser(id string) (models.User, error) {
	var u models.User
	stmt := `UPDATE users SET admin = true, updated_at = NOW() WHERE user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) DeleteUser(id string) (models.User, error) {
	var u models.User
	stmt := `DELETE FROM users  WHERE user_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}
