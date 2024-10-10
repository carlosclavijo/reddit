package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetTags() ([]models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var Tags []models.Tag
	stmt := `SELECT * FROM tags`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return Tags, err
	}
	for rows.Next() {
		var t models.Tag
		err = rows.Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return Tags, err
		}
		Tags = append(Tags, t)
	}
	return Tags, err
}

func (m *postgresDBRepo) GetTagById(tagId string) (models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Tag
	stmt := `SELECT * FROM tags WHERE tag_id = $1`
	uid, _ := uuid.FromString(tagId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

func (m *postgresDBRepo) GetTagsBySubredditId(subredditId string) ([]models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var tags []models.Tag
	stmt := `SELECT * FROM tags WHERE subreddit_id = $1`
	uid, _ := uuid.FromString(subredditId)
	rows, err := m.DB.QueryContext(ctx, stmt, uid)
	if err != nil {
		return tags, err
	}
	for rows.Next() {
		var t models.Tag
		err = rows.Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return tags, err
		}
		tags = append(tags, t)
	}
	return tags, err
}

func (m *postgresDBRepo) InsertTag(r models.Tag) (models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Tag
	stmt := `INSERT INTO tags(subreddit_id, admin_id, name, color`
	if r.IsMature {
		stmt += `, is_mature) VALUES($1, $2, $3, $4, $5) RETURNING *`
		err := m.DB.QueryRowContext(ctx, stmt, r.SubredditId, r.AdminId, r.Name, r.Color, r.IsMature).Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
		return t, err
	}
	stmt += `) VALUES($1, $2, $3, $4) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.SubredditId, r.AdminId, r.Name, r.Color).Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

func (m *postgresDBRepo) UpdateTag(tagId string, r models.Tag) (models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Tag
	stmt := `UPDATe tags SET `
	if r.Name != "" {
		stmt += `name = '` + r.Name + `', `
	} else {
		stmt += `name = name, `
	}
	if r.Color != "" {
		stmt += `color = '` + r.Color + `', `
	} else {
		stmt += `color = color, `
	}
	stmt += `is_mature = ` + strconv.FormatBool(r.IsMature) + `, updated_at = NOW() WHERE tag_id = $1 RETURNING *`
	uid, _ := uuid.FromString(tagId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}

func (m *postgresDBRepo) DeleteTag(tagId string) (models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var t models.Tag
	stmt := `DELETE FROM tags  WHERE tag_id = $1 RETURNING *`
	uid, _ := uuid.FromString(tagId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&t.TagId, &t.SubredditId, &t.AdminId, &t.Name, &t.Color, &t.IsMature, &t.CreatedAt, &t.UpdatedAt)
	return t, err
}
