package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Comment models.Comment
	err := decoder.Decode(&Comment)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	/*newComment, error := m.DB.InsertComment(Comment)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	m.App.Session.Put(r.Context(), "comment", Comment)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newComment)*/
}
