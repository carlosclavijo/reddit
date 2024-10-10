package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetPosts() ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var posts []models.Post
	stmt := `SELECT * FROM posts`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.PostId, &p.SubredditId, &p.UserId, &p.Title, &p.Description, &p.Type, &p.Nsfw, &p.Brand, &p.Votes, &p.Comments, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, p)
	}
	return posts, err
}

func (m *postgresDBRepo) GetPostById(postId string) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var p models.Post
	stmt := `SELECT * FROM posts WHERE post_id = $1`
	uid, _ := uuid.FromString(postId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&p.PostId, &p.SubredditId, &p.UserId, &p.Title, &p.Description, &p.Type, &p.Nsfw, &p.Brand, &p.Votes, &p.Comments, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func (m *postgresDBRepo) GetPostsByUserId(userId string) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var posts []models.Post
	stmt := `SELECT * FROM posts WHERE user_id = $1`
	uid, _ := uuid.FromString(userId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.PostId, &p.SubredditId, &p.UserId, &p.Title, &p.Description, &p.Type, &p.Nsfw, &p.Brand, &p.Votes, &p.Comments, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, p)
	}
	return posts, err
}

func (m *postgresDBRepo) InsertPost(r models.Post) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var p models.Post
	stmt := `INSERT INTO posts(subreddit_id, user_id, title, description`
	if r.Type != "" {
		stmt += ", type, nsfw, brand) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *"
		err := m.DB.QueryRowContext(ctx, stmt, r.SubredditId, r.UserId, r.Title, r.Description, r.Type, r.Nsfw, r.Brand).Scan(&p.PostId, &p.SubredditId, &p.UserId, &p.Title, &p.Description, &p.Type, &p.Nsfw, &p.Brand, &p.Votes, &p.Comments, &p.CreatedAt, &p.UpdatedAt)
		return p, err
	}
	stmt += ", nsfw, brand) VALUES($1, $2, $3, $4, $5, $6) RETURNING *"
	err := m.DB.QueryRowContext(ctx, stmt, r.SubredditId, r.UserId, r.Title, r.Description, r.Nsfw, r.Brand).Scan(&p.PostId, &p.SubredditId, &p.UserId, &p.Title, &p.Description, &p.Type, &p.Nsfw, &p.Brand, &p.Votes, &p.Comments, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func (m *postgresDBRepo) UpdatePost(postId string, r models.Post) (models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var p models.Post
	stmt := `UPDATE posts SET `
	if r.Title != "" {
		stmt += `title = '` + r.Title + `', `
	} else {
		stmt += `title = title, `
	}
	if r.Description != "" {
		stmt += `description = '` + r.Description + `', `
	} else {
		stmt += `description = description, `
	}
	if r.Type != "" {
		stmt += `type = '` + r.Type + `', `
	} else {
		stmt += `type = type, `
	}
	stmt += `nsfw = ` + strconv.FormatBool(r.Nsfw) + `, brand = ` + strconv.FormatBool(r.Brand) + `,  updated_at = NOW() WHERE post_id = $1 RETURNING *`
	uid, _ := uuid.FromString(postId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&p.PostId, &p.SubredditId, &p.UserId, &p.Title, &p.Description, &p.Type, &p.Nsfw, &p.Brand, &p.Votes, &p.Comments, &p.CreatedAt, &p.UpdatedAt)
	return p, err
}
