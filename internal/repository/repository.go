package repository

import "github.com/carlosclavijo/reddit/internal/models"

type DatabaseRepo interface {
	GetUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	InsertUser(r models.User) (models.User, error)
	UpdateUser(id string, r models.User) (models.User, error)
	AddUserPostKarma(id string) (models.User, error)
	AddUserCommentKarma(id string) (models.User, error)
	DeleteUser(id string) (models.User, error)

	GetSubreddits() ([]models.Subreddit, error)
	GetSubredditById(id string) (models.Subreddit, error)
	GetSubredditByUserId(id string) ([]models.Subreddit, error)
	InsertSubreddit(res models.Subreddit) (models.Subreddit, error)
	UpdateSubreddit(id string, r models.Subreddit) (models.Subreddit, error)
	DeleteSubreddit(id string) (models.Subreddit, error)

	GetSubredditsUsers() ([]models.SubredditUser, error)
	GetSubredditUserById(id string) (models.SubredditUser, error)
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

//8a1196bb-d0f9-44a5-af43-cd59b186345e
