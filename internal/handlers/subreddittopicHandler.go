package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostSubedditTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var SubredditTopic models.SubredditTopic
	err := decoder.Decode(&SubredditTopic)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertSubredditTopic(SubredditTopic)
	if error != nil {
		helpers.ServerError(w, error)

	}
	//m.App.Session.Put(r.Context(), "user", user)
}
