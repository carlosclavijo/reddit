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
	newUser, error := m.DB.InsertUser(User)
	if error != nil {
		helpers.ServerError(w, error)

	}
	m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
