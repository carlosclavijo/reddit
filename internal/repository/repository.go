package repository

import "github.com/carlosclavijo/reddit/internal/models"

type DatabaseRepo interface {
	InsertUser(res models.User) error
	InsertSubreddit(res models.Subreddit) error
	InsertSubredditUser(res models.SubredditUser) error
	InsertTopic(res models.Topic) error
	InsertSubredditTopic(res models.SubredditTopic) error
	InsertConfig(res models.Config) error
	InsertTag(res models.Tag) error
	InsertPost(res models.Post) error
	InsertPostTag(res models.PostTag) error
	InsertImage(res models.Image) error
	InsertVideo(res models.Video) error
	InsertLink(res models.Link) error
	InsertPoll(res models.Poll) error
	InsertOption(res models.Option) error
	InsertOptionUser(res models.OptionUser) error
	InsertComment(res models.Comment) error
	InsertCommentVote(res models.CommentVote) error
}
