package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetTopicsUsersList(w http.ResponseWriter, r *http.Request) {
	topicsUsers, error := m.DB.GetTopicsUsers()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(topicsUsers); i++ {
		error = m.GetTopicsAndUsers(&topicsUsers[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topicsUsers)
}

func (m *Repository) GetTopicUserById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	topicUser, error := m.DB.GetTopicUsersById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = m.GetTopicsAndUsers(&topicUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topicUser)
}

func (m *Repository) GetTopicsByUser(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	topics, error := m.DB.GetTopicsByUserId(value)
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

func (m *Repository) GetUsersByTopic(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	users, error := m.DB.GetUsersByTopicId(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
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
		helpers.ServerError(w, error)
		return
	}
	error = m.GetTopicsAndUsers(&newTopicUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
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
		return
	}
	error = m.GetTopicsAndUsers(&TopicUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = m.GetTopicsAndUsers(&TopicUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TopicUser)
}

func (m *Repository) GetTopicsAndUsers(t *models.TopicUser) error {
	var error error
	t.User, error = m.DB.GetUserById(t.UserId.String())
	if error != nil {
		return error
	}
	t.Topic, error = m.DB.GetTopicById(t.TopicId.String())
	if error != nil {
		return error
	}
	t.Topic.User, error = m.DB.GetUserById(t.Topic.UserId.String())
	return error
}
