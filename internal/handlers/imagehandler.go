package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Image models.Image
	err := decoder.Decode(&Image)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertImage(Image)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", user)
}
