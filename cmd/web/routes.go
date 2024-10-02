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
	mux.Get("/users/{userId}", handlers.Repo.GetUserById)
	mux.Post("/users", handlers.Repo.PostUser)
	mux.Put("/users/{userId}", handlers.Repo.PutUser)
	mux.Patch("/users/post/{userId}", handlers.Repo.AddPostKarma)
	mux.Patch("/users/comment/{userId}", handlers.Repo.AddCommentKarma)
	mux.Delete("/users/{userId}", handlers.Repo.DeleteUser)

	mux.Get("/subreddits", handlers.Repo.GetSubredditsList)
	mux.Get("/subreddits/users/{subredditId}", handlers.Repo.GetSubredditByUserId)
	mux.Get("/subreddits/{subredditId}", handlers.Repo.GetSubredditById)
	mux.Post("/subreddits", handlers.Repo.PostSubeddit)
	mux.Put("/subreddits/{subredditId}", handlers.Repo.PutSubreddit)
	mux.Delete("/subreddits/{subredditId}", handlers.Repo.DeleteSubreddit)

	mux.Get("/subredditsusers", handlers.Repo.GetSubredditsUsersList)
	mux.Post("/subredditusers", handlers.Repo.PostSubredditUser)

	mux.Post("/topics", handlers.Repo.PostTopic)
	mux.Post("/subreddittopics", handlers.Repo.PostSubedditTopic)
	mux.Post("/configs", handlers.Repo.PostConfig)
	mux.Post("/tags", handlers.Repo.PostTag)
	mux.Post("/posts", handlers.Repo.PostPost)
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
