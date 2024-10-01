package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostOption(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Option models.Option
	err := decoder.Decode(&Option)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	/*newOption, error := m.DB.InsertOption(Option)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	m.App.Session.Put(r.Context(), "option", Option)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newOption)*/
}
