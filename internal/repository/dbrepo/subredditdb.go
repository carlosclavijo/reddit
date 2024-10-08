package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetSubreddits() ([]models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var subreddits []models.Subreddit
	stmt := `SELECT * FROM subreddits`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return subreddits, err
	}
	for rows.Next() {
		var s models.Subreddit
		err = rows.Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return subreddits, err
		}
		subreddits = append(subreddits, s)
	}
	return subreddits, err
}

func (m *postgresDBRepo) GetSubredditById(subredditId string) (models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var s models.Subreddit
	stmt := `SELECT * FROM subreddits WHERE subreddit_id = $1`
	uid, _ := uuid.FromString(subredditId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	return s, err
}

func (m *postgresDBRepo) GetSubredditByUserId(userId string) ([]models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var subreddits []models.Subreddit
	stmt := `SELECT * FROM subreddits WHERE created_by = $1`
	uid, _ := uuid.FromString(userId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid)
	if err != nil {
		return subreddits, err
	}
	for rows.Next() {
		var s models.Subreddit
		err = rows.Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return subreddits, err
		}
		subreddits = append(subreddits, s)
	}
	return subreddits, err
}

func (m *postgresDBRepo) InsertSubreddit(r models.Subreddit) (models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var s models.Subreddit
	stmt := `INSERT INTO subreddits(name, description, created_by, privacy, is_mature`
	if r.Icon.Valid && r.Banner.Valid {
		stmt += `, icon, banner) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature, r.Icon, r.Banner).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		return s, err
	} else if r.Icon.Valid {
		stmt += `, icon) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature, r.Icon).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		return s, err
	} else if r.Banner.Valid {
		stmt += `, banner) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature, r.Banner).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		return s, err
	}
	stmt += `) VALUES($1, $2, $3, $4, $5) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	return s, err
}

func (m *postgresDBRepo) UpdateSubreddit(subredditId string, r models.Subreddit) (models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var s models.Subreddit
	stmt := `UPDATE subreddits SET `
	if r.Name != "" {
		stmt += `name = '` + r.Name + `', `
	} else {
		stmt += `name = name, `
	}
	if r.Description != "" {
		stmt += `description = '` + r.Description + `', `
	} else {
		stmt += `description = description, `
	}
	if r.Icon.String != "" {
		stmt += `icon = '` + r.Icon.String + `', `
	} else if !r.Icon.Valid {
		stmt += `icon = NULL, `
	} else {
		stmt += `icon = icon, `
	}
	if r.Banner.String != "" {
		stmt += `banner = '` + r.Banner.String + `', `
	} else if !r.Banner.Valid {
		stmt += `banner = NULL, `
	} else {
		stmt += `banner = banner, `
	}
	if r.Privacy != "" {
		stmt += `privacy = '` + r.Privacy + `', `
	} else {
		stmt += `privacy = privacy, `
	}
	stmt += `is_mature = ` + strconv.FormatBool(r.IsMature) + `, updated_at = NOW() WHERE subreddit_id = $1 RETURNING *`
	uid, _ := uuid.FromString(subredditId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	return s, err
}

func (m *postgresDBRepo) DeleteSubreddit(subredditId string) (models.Subreddit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var s models.Subreddit
	stmt := `DELETE FROM subreddits  WHERE subreddit_id = $1 RETURNING *`
	uid, _ := uuid.FromString(subredditId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	return s, err
}
