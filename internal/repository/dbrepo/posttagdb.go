package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertPostTag inserts post tags into the database
func (m *postgresDBRepo) InsertPostTag(res models.PostTag) error {
	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO posts_tags
				(post_id, tag_id)
				VALUES($1, $2)`
	_, err := m.DB.ExecContext(ctx, stmt, res.PostId, res.TagId)
	return err*/
	return nil
}
