package dbrepo

import (
	"context"
	"time"

	"github.com/carlosclavijo/reddit/internal/models"
	"github.com/gofrs/uuid"
)

func (m *postgresDBRepo) GetPostsTags() ([]models.PostTag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var posttags []models.PostTag
	stmt := `SELECT * FROM posts_tags`
	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return posttags, err
	}
	for rows.Next() {
		var pt models.PostTag
		err = rows.Scan(&pt.PostTagId, &pt.PostId, &pt.TagId, &pt.CreatedAt, &pt.UpdatedAt)
		if err != nil {
			return posttags, err
		}
		posttags = append(posttags, pt)
	}
	return posttags, err
}

func (m *postgresDBRepo) GetPostTagById(postTagId string) (models.PostTag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var pt models.PostTag
	stmt := `SELECT * FROM posts_tags WHERE post_tag_id = $1`
	uid, _ := uuid.FromString(postTagId)
	err := m.DB.QueryRowContext(ctx, stmt, uid).Scan(&pt.PostTagId, &pt.PostId, &pt.TagId, &pt.CreatedAt, &pt.UpdatedAt)
	return pt, err
}

func (m *postgresDBRepo) GetPostsByTagId(tagId string) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var posts []models.Post
	stmt := `SELECT P.* FROM posts P JOIN posts_tags PT ON P.post_id = PT.post_id JOIN tags T ON T.tag_id = PT.tag_id WHERE T.tag_id = $1`
	uid, _ := uuid.FromString(tagId)
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

func (m *postgresDBRepo) GetTagsByPostId(postId string) ([]models.Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var tags []models.Tag
	stmt := `SELECT T.* FROM tags T JOIN posts_tags PT ON T.tag_id = PT.tag_id JOIN posts P ON P.post_id = PT.post_id WHERE P.post_id = $1`
	uid, _ := uuid.FromString(postId)
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

func (m *postgresDBRepo) InsertPostTag(r models.PostTag) (models.PostTag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var pt models.PostTag
	stmt := `INSERT INTO posts_tags(post_id, tag_id) VALUES($1, $2) RETURNING *`
	err := m.DB.QueryRowContext(ctx, stmt, r.PostId, r.TagId).Scan(&pt.PostTagId, &pt.PostId, &pt.TagId, &pt.CreatedAt, &pt.UpdatedAt)
	return pt, err
}
