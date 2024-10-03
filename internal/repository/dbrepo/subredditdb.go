package dbrepo

import (
	"strconv"

	"github.com/carlosclavijo/reddit/internal/models"
)

// GetReddits get the list of all reddits from the database
func (m *postgresDBRepo) GetSubreddits() ([]models.Subreddit, error) {
	var subreddits []models.Subreddit
	stmt := `SELECT * FROM subreddits`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return subreddits, err
	}
	for rows.Next() {
		var s models.Subreddit
		err = rows.Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return subreddits, err
		}
		s.User, err = m.GetUserById(s.CreatedBy.String())
		if err != nil {
			return subreddits, err
		}
		subreddits = append(subreddits, s)
	}
	return subreddits, err
}

// GetSubredditById gets the subreddit with their uuid
func (m *postgresDBRepo) GetSubredditById(id string) (models.Subreddit, error) {
	var s models.Subreddit
	stmt := `SELECT * FROM subreddits WHERE subreddit_id = '` + id + `'`
	err := m.DB.QueryRow(stmt).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return s, err
	}
	s.User, err = m.GetUserById(s.CreatedBy.String())
	return s, err
}

// GetSubredditByUserId gets all the subreddits created by userId
func (m *postgresDBRepo) GetSubredditByUserId(id string) ([]models.Subreddit, error) {
	var subreddits []models.Subreddit
	stmt := `SELECT * FROM subreddits WHERE created_by = '` + id + `'`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return subreddits, err
	}
	for rows.Next() {
		var s models.Subreddit
		err = rows.Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return subreddits, err
		}
		s.User, err = m.GetUserById(s.CreatedBy.String())
		subreddits = append(subreddits, s)
	}
	return subreddits, err
}

// InsertSubreddit inserts subreddits into the database
func (m *postgresDBRepo) InsertSubreddit(r models.Subreddit) (models.Subreddit, error) {
	var s models.Subreddit
	stmt := `INSERT INTO subreddits(name, description, created_by, privacy, is_mature`
	if r.Icon.Valid && r.Banner.Valid {
		stmt += `, icon, banner) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`
		err := m.DB.QueryRow(stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature, r.Icon, r.Banner).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		return s, err
	} else if r.Icon.Valid {
		stmt += `, icon) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`
		err := m.DB.QueryRow(stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature, r.Icon).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		return s, err
	} else if r.Banner.Valid {
		stmt += `, banner) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`
		err := m.DB.QueryRow(stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature, r.Banner).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
		return s, err
	}
	stmt += `) VALUES($1, $2, $3, $4, $5) RETURNING *`
	err := m.DB.QueryRow(stmt, r.Name, r.Description, r.CreatedBy, r.Privacy, r.IsMature).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return s, err
	}
	s.User, err = m.GetUserById(r.CreatedBy.String())
	return s, err
}

// UpdateSubreddit updates subreddit information
func (m *postgresDBRepo) UpdateSubreddit(id string, r models.Subreddit) (models.Subreddit, error) {
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
	stmt += `is_mature = ` + strconv.FormatBool(r.IsMature) + `, updated_at = NOW() WHERE subreddit_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	return s, err
}

// DeleteSubreddit deleters the subreddit from the database
func (m *postgresDBRepo) DeleteSubreddit(id string) (models.Subreddit, error) {
	var s models.Subreddit
	stmt := `DELETE FROM subreddits  WHERE subreddit_id = '` + id + `' RETURNING *`
	err := m.DB.QueryRow(stmt).Scan(&s.SubredditId, &s.Name, &s.Description, &s.CreatedBy, &s.Icon, &s.Banner, &s.Privacy, &s.IsMature, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return s, err
	}
	s.User, err = m.GetUserById(s.CreatedBy.String())
	return s, err
}
