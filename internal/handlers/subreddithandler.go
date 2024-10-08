package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetSubredditsList(w http.ResponseWriter, r *http.Request) {
	subreddits, error := m.DB.GetSubreddits()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(subreddits); i++ {
		subreddits[i].User, error = m.DB.GetUserById(subreddits[i].CreatedBy.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
		subreddits[i].User.Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subreddits)
}

func (m *Repository) GetSubredditById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	subreddit, error := m.DB.GetSubredditById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	subreddit.User, error = m.DB.GetUserById(subreddit.CreatedBy.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	subreddit.User.Password = "restricted"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subreddit)
}

func (m *Repository) GetSubredditByUserId(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	subreddits, error := m.DB.GetSubredditByUserId(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(subreddits); i++ {
		subreddits[i].User, error = m.DB.GetUserById(subreddits[i].CreatedBy.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
		subreddits[i].User.Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subreddits)
}

func (m *Repository) PostSubeddit(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Subreddit models.Subreddit
	err := decoder.Decode(&Subreddit)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newSubreddit, error := m.DB.InsertSubreddit(Subreddit)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newSubreddit.User, error = m.DB.GetUserById(newSubreddit.CreatedBy.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newSubreddit.User.Password = "restricted"
	//m.App.Session.Put(r.Context(), "subreddit", Subreddit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubreddit)
}

func (m *Repository) PutSubreddit(w http.ResponseWriter, r *http.Request) {
	var Subreddit models.Subreddit
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&Subreddit)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newSubreddit, error := m.DB.UpdateSubreddit(value, Subreddit)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newSubreddit.User, error = m.DB.GetUserById(newSubreddit.CreatedBy.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newSubreddit.User.Password = "restricted"
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubreddit)
}

func (m *Repository) DeleteSubreddit(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	Subreddit, error := m.DB.DeleteSubreddit(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	Subreddit.User, error = m.DB.GetUserById(Subreddit.CreatedBy.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	Subreddit.User.Password = "restricted"
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Subreddit)
}
