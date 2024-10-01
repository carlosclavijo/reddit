package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Tag models.Tag
	err := decoder.Decode(&Tag)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	/*newTag, error := m.DB.InsertTag(Tag)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	m.App.Session.Put(r.Context(), "tag", Tag)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTag)*/
}
