package dbrepo

import (
	"context"
	"strconv"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

// GetConfigs get the list of all configs from the database
func (m *postgresDBRepo) GetConfigs() ([]models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var configs []models.Config
	stmt := `SELECT * FROM configs`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return configs, err
	}
	for rows.Next() {
		var c models.Config
		err = rows.Scan(&c.ConfigId, &c.SubredditId, &c.AdminId, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return configs, err
		}
		configs = append(configs, c)
	}
	return configs, err
}

// GetConfigById gets the config with its uuid
func (m *postgresDBRepo) GetConfigById(configId string) (models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var c models.Config
	stmt := `SELECT * FROM configs WHERE config_id = $1`
	uid, _ := uuid.FromString(configId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&c.ConfigId, &c.SubredditId, &c.AdminId, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}

// InsertConfig inserts configurations into the database
func (m *postgresDBRepo) InsertConfig(res models.Config) (models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var c models.Config
	stmt := `INSERT INTO configs (subreddit_id, admin_config) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, res.SubredditId, res.AdminId).Scan(&c.ConfigId, &c.SubredditId, &c.AdminId, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}

// UpdateConfig updates topic information
func (m *postgresDBRepo) UpdateConfig(configId string, r models.Config) (models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var c models.Config
	stmt := `UPDATE configs SET `
	if r.AdminId.String() != "" {
		stmt += `admin_config = '` + r.AdminId.String() + `', `
	} else {
		stmt += `admin_config = admin_config, `
	}
	stmt += `is_available = ` + strconv.FormatBool(r.IsAvailable) + `, ` +
		`is_locked = ` + strconv.FormatBool(r.IsLocked) + `, ` +
		`text_available = ` + strconv.FormatBool(r.TextAvailable) + `, ` +
		`images_available = ` + strconv.FormatBool(r.ImageAvailable) + `, ` +
		`video_available = ` + strconv.FormatBool(r.VideoAvailable) + `, ` +
		`link_available = ` + strconv.FormatBool(r.LinkAvailable) + `, ` +
		`poll_available = ` + strconv.FormatBool(r.PollAvailable) + `, ` +
		`updated_at = NOW() WHERE config_id = $1 RETURNING *`
	uid, _ := uuid.FromString(configId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&c.ConfigId, &c.SubredditId, &c.AdminId, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}

// DeleteConfig deletes the config
func (m *postgresDBRepo) DeleteConfig(configId string) (models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var c models.Config
	stmt := `DELETE FROM configs WHERE config_id = $1 RETURNING *`
	uid, _ := uuid.FromString(configId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&c.ConfigId, &c.SubredditId, &c.AdminId, &c.IsAvailable, &c.IsLocked, &c.TextAvailable, &c.ImageAvailable, &c.VideoAvailable, &c.LinkAvailable, &c.PollAvailable, &c.CreatedAt, &c.UpdatedAt)
	return c, err
}
