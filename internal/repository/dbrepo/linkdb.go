package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertLink inserts links into the database
func (m *postgresDBRepo) InsertLink(res models.Link) error {
	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO links
				(post_id, link)
				VALUES($1, $2)`
	_, err := m.DB.ExecContext(ctx, stmt, res.PostId, res.Link)
	return err*/
	return nil
}
