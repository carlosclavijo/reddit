package main

import (
	"net/http"

	"github.com/carlosclavijo/reddit/internal/config"
	"github.com/carlosclavijo/reddit/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HelloWorld)

	mux.Get("/users", handlers.Repo.GetUsersList)
	mux.Get("/users/admins", handlers.Repo.GetUsersAdminList)
	mux.Get("/users/{userId}", handlers.Repo.GetUserById)
	mux.Post("/users", handlers.Repo.PostUser)
	mux.Put("/users/{userId}", handlers.Repo.PutUser)
	mux.Patch("/users/post/plus/{userId}", handlers.Repo.PatchPlusPostKarma)
	mux.Patch("/users/post/less/{userId}", handlers.Repo.PatchLessPostKarma)
	mux.Patch("/users/comment/plus/{userId}", handlers.Repo.PatchPlusCommentKarma)
	mux.Patch("/users/comment/less/{userId}", handlers.Repo.PatchLessCommentKarma)
	mux.Patch("/users/admin/{userId}", handlers.Repo.PatchAdmin)
	mux.Delete("/users/{userId}", handlers.Repo.DeleteUser)

	mux.Get("/topics", handlers.Repo.GetTopicsList)
	mux.Get("/topics/{topicId}", handlers.Repo.GetTopicById)
	mux.Get("/topics/sub/{topicId}", handlers.Repo.GetSubtopics)
	mux.Get("/topics/parents", handlers.Repo.GetParentsTopicsList)
	mux.Get("/topics/creator/{userId}", handlers.Repo.GetTopicsByCreatorId)
	mux.Post("/topics", handlers.Repo.PostTopic)
	mux.Put("/topics/{topicId}", handlers.Repo.PutTopic)
	mux.Delete("/topics/{topicId}", handlers.Repo.DeleteTopic)

	mux.Get("/topicsusers", handlers.Repo.GetTopicsUsersList)
	mux.Get("/topicsusers/{topicId}", handlers.Repo.GetTopicUserById)
	mux.Get("/topicsusers/topics/{userId}", handlers.Repo.GetTopicsByUser)
	mux.Get("/topicsusers/users/{topicId}", handlers.Repo.GetUsersByTopic)
	mux.Post("/topicsusers", handlers.Repo.PostTopicUser)
	mux.Delete("/topicsusers/{topicUserId}", handlers.Repo.DeleteTopicUser)

	mux.Get("/subreddits", handlers.Repo.GetSubredditsList)
	mux.Get("/subreddits/users/{subredditId}", handlers.Repo.GetSubredditByUserId)
	mux.Get("/subreddits/{subredditId}", handlers.Repo.GetSubredditById)
	mux.Post("/subreddits", handlers.Repo.PostSubeddit)
	mux.Put("/subreddits/{subredditId}", handlers.Repo.PutSubreddit)
	mux.Delete("/subreddits/{subredditId}", handlers.Repo.DeleteSubreddit)

	mux.Get("/configs", handlers.Repo.GetConfigsList)
	mux.Get("/configs/{configId}", handlers.Repo.GetConfigById)
	mux.Post("/configs", handlers.Repo.PostConfig)
	mux.Put("/configs/{configId}", handlers.Repo.PutConfig)
	mux.Delete("/configs/{cofigId}", handlers.Repo.DeleteConfig)

	mux.Get("/subredditsusers", handlers.Repo.GetSubredditsUsersList)
	mux.Get("/subredditsusers/{subredditUserId}", handlers.Repo.GetSubredditUserById)
	mux.Get("/subredditsusers/members/{subredditId}", handlers.Repo.GetSubredditMembers)
	mux.Get("/subredditsusers/members/{role}/{subredditId}", handlers.Repo.GetSubredditMembersByRole)
	mux.Post("/subredditsusers", handlers.Repo.PostSubredditUser)
	mux.Put("/subredditusers/{subredditUserId}", handlers.Repo.PutSubredditUser)
	mux.Delete("/subredditusers/{subredditUserId}", handlers.Repo.DeleteSubredditUser)

	mux.Get("/subreddittopics", handlers.Repo.GetSubredditsTopicsList)
	mux.Get("/subreddittopics/{subredditTopicId}", handlers.Repo.GetSubredditTopicById)
	mux.Get("/subreddittopics/subreddits/{topicId}", handlers.Repo.GetSubredditsByTopic)
	mux.Get("/subreddittopics/topics/{subredditId}", handlers.Repo.GetTopicsBySubreddit)
	mux.Post("/subreddittopics", handlers.Repo.PostSubedditTopic)
	mux.Delete("/subreddittopics/{subredditTopicId}", handlers.Repo.DeleteSubredditTopic)

	mux.Get("/tags", handlers.Repo.GetTagsList)
	mux.Get("/tags/{tagId}", handlers.Repo.GetTagById)
	mux.Get("/tags/subreddits/{subredditId}", handlers.Repo.GetTagsBySubreddit)
	mux.Post("/tags", handlers.Repo.PostTag)
	mux.Put("/tags/{tagId}", handlers.Repo.PutTag)
	mux.Delete("/tags/{tagId}", handlers.Repo.DeleteTag)

	mux.Get("/posts", handlers.Repo.GetPostsList)
	mux.Get("/posts/{postId}", handlers.Repo.GetPostById)
	mux.Get("/posts/user/{userId}", handlers.Repo.GetPostsByUserId)
	mux.Post("/posts", handlers.Repo.PostPost)
	mux.Put("/posts/{postId}", handlers.Repo.PutPost)

	mux.Post("/posttags", handlers.Repo.PostPostTag)
	mux.Post("/images", handlers.Repo.PostImage)
	mux.Post("/videos", handlers.Repo.PostVideo)
	mux.Post("/links", handlers.Repo.PostLink)
	mux.Post("/polls", handlers.Repo.PostPoll)
	mux.Post("/options", handlers.Repo.PostOption)
	mux.Post("/optionusers", handlers.Repo.PostOptionUser)
	mux.Post("/comments", handlers.Repo.PostComment)
	mux.Post("/commentvotes", handlers.Repo.PostCommentVote)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
