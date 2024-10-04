package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetTopicsUsersList(w http.ResponseWriter, r *http.Request) {
	topicsUsers, error := m.DB.GetTopicsUsers()
	if error != nil {
		helpers.ServerError(w, error)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topicsUsers)
}

func (m *Repository) GetTopicUserById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	topicUser, error := m.DB.GetTopicUsersById(value)
	if error != nil {
		helpers.ServerError(w, error)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topicUser)
}

func (m *Repository) PostTopicUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var TopicUser models.TopicUser
	err := decoder.Decode(&TopicUser)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newTopicUser, error := m.DB.InsertTopicUser(TopicUser)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "topic", Topic)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTopicUser)
}

func (m *Repository) DeleteTopicUser(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	TopicUser, error := m.DB.DeleteTopicUser(value)
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TopicUser)
}
