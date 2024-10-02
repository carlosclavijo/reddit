package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetSubredditsUsersList(w http.ResponseWriter, r *http.Request) {
	subredditsusers, error := m.DB.GetSubredditsUsers()
	if error != nil {
		helpers.ServerError(w, error)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subredditsusers)
}

func (m *Repository) PostSubredditUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var SubredditUser models.SubredditUser
	err := decoder.Decode(&SubredditUser)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newSubredditUser, error := m.DB.InsertSubredditUser(SubredditUser)
	if error != nil {
		helpers.ServerError(w, error)
	}
	m.App.Session.Put(r.Context(), "subreddituser", SubredditUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubredditUser)
}
