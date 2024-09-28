package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostPostTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var PostTag models.PostTag
	err := decoder.Decode(&PostTag)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertPostTag(PostTag)
	if error != nil {
		helpers.ServerError(w, error)

	}
	//m.App.Session.Put(r.Context(), "user", user)
}
