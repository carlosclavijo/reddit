package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

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

	}
	m.App.Session.Put(r.Context(), "subreddit", Subreddit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubreddit)
}
