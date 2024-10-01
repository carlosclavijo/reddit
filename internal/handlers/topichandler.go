package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Topic models.Topic
	err := decoder.Decode(&Topic)
	if err != nil {
		log.Println(err)
		helpers.ServerError(w, err)
		return
	}
	newTopic, error := m.DB.InsertTopic(Topic)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)

	}
	m.App.Session.Put(r.Context(), "topic", Topic)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTopic)
}
