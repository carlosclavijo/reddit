package repository

import "github.com/carlosclavijo/reddit/internal/models"

type DatabaseRepo interface {
	GetUsers() ([]models.User, error)
	GetUserById(userId string) (models.User, error)
	GetUsersAdmins() ([]models.User, error)
	InsertUser(r models.User) (models.User, error)
	UpdateUser(userId string, r models.User) (models.User, error)
	PlusUserPostKarma(userId string) (models.User, error)
	LessUserPostKarma(userId string) (models.User, error)
	PlusUserCommentKarma(userId string) (models.User, error)
	LessUserCommentKarma(userId string) (models.User, error)
	AdminUser(userId string) (models.User, error)
	DeleteUser(userId string) (models.User, error)

	GetTopics() ([]models.Topic, error)
	GetTopicById(topicId string) (models.Topic, error)
	GetSubTopics(topicId string) ([]models.Topic, error)
	GetParentsTopics() ([]models.Topic, error)
	GetTopicsByCreatorId(userId string) ([]models.Topic, error)
	InsertTopic(r models.Topic) (models.Topic, error)
	UpdateTopic(topicId string, r models.Topic) (models.Topic, error)
	DeleteTopic(topicId string) (models.Topic, error)

	GetTopicsUsers() ([]models.TopicUser, error)
	GetTopicUsersById(topicUserId string) (models.TopicUser, error)
	GetTopicsByUserId(userId string) ([]models.Topic, error)
	GetUsersByTopicId(topicId string) ([]models.User, error)
	InsertTopicUser(r models.TopicUser) (models.TopicUser, error)
	DeleteTopicUser(topicUserId string) (models.TopicUser, error)

	GetSubreddits() ([]models.Subreddit, error)
	GetSubredditById(subredditId string) (models.Subreddit, error)
	GetSubredditByUserId(userId string) ([]models.Subreddit, error)
	InsertSubreddit(r models.Subreddit) (models.Subreddit, error)
	UpdateSubreddit(subredditId string, r models.Subreddit) (models.Subreddit, error)
	DeleteSubreddit(subredditId string) (models.Subreddit, error)

	GetConfigs() ([]models.Config, error)
	GetConfigById(configId string) (models.Config, error)
	InsertConfig(res models.Config) (models.Config, error)
	UpdateConfig(configId string, r models.Config) (models.Config, error)
	DeleteConfig(configId string) (models.Config, error)

	GetSubredditsUsers() ([]models.SubredditUser, error)
	GetSubredditUserById(subredditUserId string) (models.SubredditUser, error)
	GetSubredditMembers(subredditUserId string) ([]models.User, error)
	GetSubredditMembersByRole(subredditUserId string, role string) ([]models.User, error)
	InsertSubredditUser(r models.SubredditUser) (models.SubredditUser, error)
	UpdateSubredditUser(subredditUserId string, r models.SubredditUser) (models.SubredditUser, error)
	DeleteSubredditUser(subredditUserId string) (models.SubredditUser, error)

	GetSubredditsTopics() ([]models.SubredditTopic, error)
	GetSubredditsTopicById(subredditTopicId string) (models.SubredditTopic, error)
	GetSubredditsByTopicId(topicId string) ([]models.Subreddit, error)
	GetTopicsBySubredditId(subredditId string) ([]models.Topic, error)
	InsertSubredditTopic(res models.SubredditTopic) (models.SubredditTopic, error)
	DeleteSubredditTopic(subredditTopicId string) (models.SubredditTopic, error)

	GetTags() ([]models.Tag, error)
	GetTagById(tagId string) (models.Tag, error)
	GetTagsBySubredditId(subredditId string) ([]models.Tag, error)
	InsertTag(r models.Tag) (models.Tag, error)
	UpdateTag(tagId string, r models.Tag) (models.Tag, error)
	DeleteTag(tagId string) (models.Tag, error)

	GetPosts() ([]models.Post, error)
	GetPostById(postId string) (models.Post, error)
	GetPostsByUserId(userId string) ([]models.Post, error)
	InsertPost(r models.Post) (models.Post, error)
	UpdatePost(postId string, r models.Post) (models.Post, error)
	ChangeNsfw(r models.Post) (models.Post, error)
	//ChangeNsfw(postId string) (models.Post, error)
	//DeletePost(postId string) (models.Post, error)

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
