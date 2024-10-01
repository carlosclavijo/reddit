package dbrepo

import (
	"log"

	"github.com/carlosclavijo/reddit/internal/models"
)

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
	stmt = `SELECT user_id, username, email, post_karma, comment_karma, account_available, profile_pic, created_at, updated_at FROM users WHERE user_id = '` + r.CreatedBy.String() + `'`
	log.Println(stmt)
	err = m.DB.QueryRow(stmt).Scan(&s.User.UserId, &s.User.Username, &s.User.Email, &s.User.PostKarma, &s.User.CommentKarma, &s.User.AccountAvailable, &s.User.ProfilePic, &s.User.CreatedAt, &s.User.UpdatedAt)
	return s, err
}
