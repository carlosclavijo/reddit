package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetPostsList(w http.ResponseWriter, r *http.Request) {
	posts, error := m.DB.GetPosts()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(posts); i++ {
		error = getSubredditsAndUsersByPost(&posts[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (m *Repository) GetPostById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	post, error := m.DB.GetPostById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByPost(&post)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (m *Repository) GetPostsByUserId(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	posts, error := m.DB.GetPostsByUserId(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(posts); i++ {
		error = getSubredditsAndUsersByPost(&posts[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (m *Repository) PostPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Post models.Post
	err := decoder.Decode(&Post)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newPost, error := m.DB.InsertPost(Post)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByPost(&newPost)
	if error != nil {
		helpers.ServerError(w, err)
		return
	}
	//m.App.Session.Put(r.Context(), "post", Post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func (m *Repository) PutPost(w http.ResponseWriter, r *http.Request) {
	var Post models.Post
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&Post)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newPost, error := m.DB.UpdatePost(value, Post)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByPost(&newPost)
	if error != nil {
		helpers.ServerError(w, err)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func (m *Repository) PatchPostNsfw(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	post, err := m.DB.GetPostById(value)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newPost, error := m.DB.ChangeNsfw(value, post.Nsfw)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByPost(&newPost)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func (m *Repository) PatchPostBrand(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	post, err := m.DB.GetPostById(value)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newPost, error := m.DB.ChangeBrand(value, post.Brand)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByPost(&newPost)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func (m *Repository) DeletePost(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	Post, error := m.DB.DeletePost(value)
	if error != nil {

		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByPost(&Post)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Post)
}

func getSubredditsAndUsersByPost(p *models.Post) error {
	var error error
	p.User, error = Repo.DB.GetUserById(p.UserId.String())
	if error != nil {
		return error
	}
	p.User.Password = "restricted"
	p.Subreddit, error = Repo.DB.GetSubredditById(p.SubredditId.String())
	if error != nil {
		return error
	}
	p.Subreddit.User, error = Repo.DB.GetUserById(p.Subreddit.CreatedBy.String())
	p.Subreddit.User.Password = "restricted"
	return error
}
