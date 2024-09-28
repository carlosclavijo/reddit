package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var User models.User
	err := decoder.Decode(&User)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertUser(User)
	if error != nil {
		helpers.ServerError(w, error)

	}
	//m.App.Session.Put(r.Context(), "user", user)
}
