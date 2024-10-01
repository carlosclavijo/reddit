package dbrepo

import (
	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertOption inserts optionf of polls into the database
func (m *postgresDBRepo) InsertOption(res models.Option) error {
	/*ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO options
				(poll_id, value)
				VALUES($1, $2)`
	_, err := m.DB.ExecContext(ctx, stmt, res.PollId, res.Value)
	return err*/
	return nil
}
