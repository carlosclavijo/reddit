package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
)

// InsertCommentVote inserts comment votes into the database
func (m *postgresDBRepo) InsertCommentVote(res models.CommentVote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO comments_vote
				(comment_id, user_id, vote)
				VALUES($1, $2, $3)`
	_, err := m.DB.ExecContext(ctx, stmt, res.CommentId, res.UserId, res.Vote)
	return err
}
