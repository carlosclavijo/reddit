package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetTopicsList(w http.ResponseWriter, r *http.Request) {
	topics, error := m.DB.GetTopics()
	if error != nil {
		helpers.ServerError(w, error)
	}
	for i := 0; i < len(topics); i++ {
		topics[i].User, error = m.DB.GetUserById(topics[i].UserId.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topics)
}

func (m *Repository) GetTopicById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	topic, error := m.DB.GetTopicById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	topic.User, error = m.DB.GetUserById(topic.UserId.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topic)
}

func (m *Repository) GetSubtopics(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	topics, error := m.DB.GetSubTopics(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(topics); i++ {
		topics[i].User, error = m.DB.GetUserById(topics[i].UserId.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topics)
}

func (m *Repository) GetParentsTopicsList(w http.ResponseWriter, r *http.Request) {
	topics, error := m.DB.GetParentsTopics()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(topics); i++ {
		topics[i].User, error = m.DB.GetUserById(topics[i].UserId.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topics)
}

func (m *Repository) PostTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Topic models.Topic
	err := decoder.Decode(&Topic)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	user, err := m.DB.GetUserById(Topic.UserId.String())
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	if !user.Admin {
		helpers.ServerError(w, errors.New("you can't add a topic because you're not an admin"))
		return
	}
	newTopic, error := m.DB.InsertTopic(Topic)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newTopic.User, error = m.DB.GetUserById(Topic.UserId.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "topic", Topic)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTopic)
}

func (m *Repository) PutTopic(w http.ResponseWriter, r *http.Request) {
	var Topic models.Topic
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&Topic)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newTopic, error := m.DB.UpdateTopic(value, Topic)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newTopic.User, error = m.DB.GetUserById(newTopic.UserId.String())
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTopic)
}

func (m *Repository) DeleteTopic(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	Topic, error := m.DB.DeleteTopic(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	Topic.User, error = m.DB.GetUserById(Topic.UserId.String())
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Topic)
}
