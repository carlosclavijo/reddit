package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetPostsTags(w http.ResponseWriter, r *http.Request) {
	postsTags, error := m.DB.GetPostsTags()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(postsTags); i++ {
		error = getPostsAndTagsByPostTag(&postsTags[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postsTags)
}

func (m *Repository) GetPostTagById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	postTag, error := m.DB.GetPostTagById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getPostsAndTagsByPostTag(&postTag)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postTag)
}

func (m *Repository) GetPostsByTagId(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	posts, error := m.DB.GetPostsByTagId(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(posts); i++ {
		posts[i].User, error = m.DB.GetUserById(posts[i].UserId.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
		posts[i].User.Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (m *Repository) GetTagsByPostId(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	tags, error := m.DB.GetTagsByPostId(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(tags); i++ {
		tags[i].Admin, error = m.DB.GetUserById(tags[i].AdminId.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
		tags[i].Admin.Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (m *Repository) PostPostTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var PostTag models.PostTag
	err := decoder.Decode(&PostTag)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newPostTag, error := m.DB.InsertPostTag(PostTag)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getPostsAndTagsByPostTag(&newPostTag)
	if error != nil {
		log.Print(error)
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "posttag", PostTag)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPostTag)
}

func getPostsAndTagsByPostTag(pt *models.PostTag) error {
	var error error
	pt.Post, error = Repo.DB.GetPostById(pt.PostId.String())
	if error != nil {
		return error
	}
	pt.Post.User, error = Repo.DB.GetUserById(pt.Post.UserId.String())
	if error != nil {
		return error
	}
	pt.Post.User.Password = "restricted"
	pt.Post.Subreddit, error = Repo.DB.GetSubredditById(pt.Post.SubredditId.String())
	if error != nil {
		return error
	}
	pt.Post.Subreddit.User, error = Repo.DB.GetUserById(pt.Post.Subreddit.CreatedBy.String())
	if error != nil {
		return error
	}
	pt.Post.Subreddit.User.Password = "restricted"
	pt.Tag, error = Repo.DB.GetTagById(pt.TagId.String())
	if error != nil {
		return error
	}
	pt.Tag.Admin, error = Repo.DB.GetUserById(pt.Tag.AdminId.String())
	if error != nil {
		return error
	}
	pt.Tag.Admin.Password = "restricted"
	pt.Tag.Subreddit, error = Repo.DB.GetSubredditById(pt.Tag.SubredditId.String())
	if error != nil {
		return error
	}
	pt.Tag.Subreddit.User, error = Repo.DB.GetUserById(pt.Post.Subreddit.CreatedBy.String())
	pt.Tag.Subreddit.User.Password = "restricted"
	return error
}
