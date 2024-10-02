package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetUsersList(w http.ResponseWriter, r *http.Request) {
	users, error := m.DB.GetUsers()
	if error != nil {
		helpers.ServerError(w, error)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (m *Repository) GetUserById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	user, error := m.DB.GetUserById(value)
	if error != nil {
		helpers.ServerError(w, error)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (m *Repository) PostUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newUser, error := m.DB.InsertUser(User)
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func (m *Repository) PutUser(w http.ResponseWriter, r *http.Request) {
	var User models.User
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newUser, error := m.DB.UpdateUser(value, User)
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func (m *Repository) AddPostKarma(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	User, error := m.DB.AddUserPostKarma(value)
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}

func (m *Repository) AddCommentKarma(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	User, error := m.DB.AddUserCommentKarma(value)
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}

func (m *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	User, error := m.DB.DeleteUser(value)
	if error != nil {
		helpers.ServerError(w, error)
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}
