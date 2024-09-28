package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostOptionUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var OptionUser models.OptionUser
	err := decoder.Decode(&OptionUser)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newOptionUser, error := m.DB.InsertOptionUser(OptionUser)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
		return
	}
	m.App.Session.Put(r.Context(), "optionuser", OptionUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOptionUser)
}
