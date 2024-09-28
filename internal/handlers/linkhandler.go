package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostLink(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Link models.Link
	err := decoder.Decode(&Link)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertLink(Link)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", user)
}
