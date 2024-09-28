package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Topic models.Topic
	err := decoder.Decode(&Topic)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertTopic(Topic)
	if error != nil {
		helpers.ServerError(w, error)

	}
	//m.App.Session.Put(r.Context(), "user", user)
}
