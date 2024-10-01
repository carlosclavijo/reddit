package repository

import "github.com/carlosclavijo/reddit/internal/models"

type DatabaseRepo interface {
	InsertUser(res models.User) (models.User, error)
	GetUser(id string) (models.User, error)

	InsertSubreddit(res models.Subreddit) (models.Subreddit, error)
	InsertSubredditUser(res models.SubredditUser) (models.SubredditUser, error)
	InsertTopic(res models.Topic) (models.Topic, error)
	InsertSubredditTopic(res models.SubredditTopic) (models.SubredditTopic, error)
	InsertConfig(res models.Config) (models.Config, error)
	//InsertTag(res models.Tag) (models.Tag, error)
	/*InsertPost(res models.Post) (models.Post, error)
	InsertPostTag(res models.PostTag) (models.PostTag, error)
	InsertImage(res models.Image) (models.Image, error)
	InsertVideo(res models.Video) (models.Video, error)
	InsertLink(res models.Link) (models.Link, error)
	InsertPoll(res models.Poll) (models.Poll, error)
	InsertOption(res models.Option) (models.Option, error)
	InsertOptionUser(res models.OptionUser) (models.OptionUser, error)
	InsertComment(res models.Comment) (models.Comment, error)
	InsertCommentVote(res models.CommentVote) (models.CommentVote, error)*/
}
