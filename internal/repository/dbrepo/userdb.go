package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

// GetUsers get the list of all users from the database
func (m *postgresDBRepo) Users() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []models.User
	stmt := `SELECT * FROM users`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return users, err
		}
		u.Password = "restricted"
		users = append(users, u)
	}
	return users, err
}

// GetUser gets user by id from the database
func (m *postgresDBRepo) GetUserById(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `SELECT * FROM users WHERE user_id = $1`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	u.Password = "restricted"
	return u, err

}

// InsertUser inserts users into the database
func (m *postgresDBRepo) InsertUser(r models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `INSERT INTO users (username, email, password`
	if r.ProfilePic.Valid && r.Admin {
		stmt += `, profile_pic, admin) VALUES($1, $2, $3, $4, $5) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.Username, r.Email, r.Password, r.ProfilePic, r.Admin).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		return u, err
	} else if r.ProfilePic.Valid && !r.Admin {
		stmt += `, profile_pic) VALUES($1, $2, $3, $4) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.Username, r.Email, r.Password, r.ProfilePic).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		return u, err
	} else if !r.ProfilePic.Valid && r.Admin {
		stmt += `, admin) VALUES($1, $2, $3, $4) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.Username, r.Email, r.Password, r.Admin).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
		return u, err
	}
	stmt += `) VALUES($1, $2, $3) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.Username, r.Email, r.Password).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	u.Password = "restricted"
	return u, err
}

// UpdateUser updates user information
func (m *postgresDBRepo) UpdateUser(id string, r models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
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
		stmt += `profile_pic = NULL`
	} else {
		stmt += `profile_pic = profile_pic`
	}
	stmt += `, admin = ` + strconv.FormatBool(r.Admin) + `, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AddUserPostKarma(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `UPDATE users SET post_karma = post_karma + 1, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AddUserCommentKarma(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `UPDATE users SET comment_karma = comment_karma + 1, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AdminUser(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `UPDATE users SET admin = true, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) DeleteUser(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `DELETE FROM users  WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(id)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}
