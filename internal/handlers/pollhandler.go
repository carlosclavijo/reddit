package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostPoll(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Poll models.Poll
	err := decoder.Decode(&Poll)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	error := m.DB.InsertPoll(Poll)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", user)
}
