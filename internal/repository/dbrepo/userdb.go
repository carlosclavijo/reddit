package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetUsers() ([]models.User, error) {
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
		users = append(users, u)
	}
	return users, err
}

func (m *postgresDBRepo) GetUsersAdmins() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var users []models.User
	stmt := `SELECT * FROM users WHERE admin = true`
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
		users = append(users, u)
	}
	return users, err
}

func (m *postgresDBRepo) GetUserById(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `SELECT * FROM users WHERE user_id = 
	$1`
	uid, _ := uuid.FromString(userId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

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
	return u, err
}

func (m *postgresDBRepo) UpdateUser(userId string, r models.User) (models.User, error) {
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
	uid, _ := uuid.FromString(userId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AddUserPostKarma(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `UPDATE users SET post_karma = post_karma + 1, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(userId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AddUserCommentKarma(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `UPDATE users SET comment_karma = comment_karma + 1, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(userId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) AdminUser(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `UPDATE users SET admin = true, updated_at = NOW() WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(userId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func (m *postgresDBRepo) DeleteUser(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var u models.User
	stmt := `DELETE FROM users  WHERE user_id = $1 RETURNING *`
	uid, _ := uuid.FromString(userId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&u.UserId, &u.Username, &u.Email, &u.Password, &u.PostKarma, &u.CommentKarma, &u.AccountAvailable, &u.ProfilePic, &u.Admin, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}
