package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Post models.Post
	err := decoder.Decode(&Post)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	/*newPoll, error := m.DB.InsertPost(Post)
	if error != nil {
		helpers.ServerError(w, error)

	}
	m.App.Session.Put(r.Context(), "post", Post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPoll)*/
}
