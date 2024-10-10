package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetSubredditsUsersList(w http.ResponseWriter, r *http.Request) {
	subredditsusers, error := m.DB.GetSubredditsUsers()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(subredditsusers); i++ {
		error = getSubredditsAndUsersBySubredditUser(&subredditsusers[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subredditsusers)
}

func (m *Repository) GetSubredditUserById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	subreddit, error := m.DB.GetSubredditUserById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersBySubredditUser(&subreddit)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subreddit)
}

func (m *Repository) GetSubredditMembers(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	users, error := m.DB.GetSubredditMembers(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(users); i++ {
		users[i].Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (m *Repository) GetSubredditMembersByRole(w http.ResponseWriter, r *http.Request) {
	role := strings.Split(r.URL.Path, "/")[3]
	value := strings.Split(r.URL.Path, "/")[4]
	users, error := m.DB.GetSubredditMembersByRole(value, role)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (m *Repository) PostSubredditUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var SubredditUser models.SubredditUser
	err := decoder.Decode(&SubredditUser)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newSubredditUser, error := m.DB.InsertSubredditUser(SubredditUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersBySubredditUser(&newSubredditUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "subreddituser", SubredditUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubredditUser)
}

func (m *Repository) PutSubredditUser(w http.ResponseWriter, r *http.Request) {
	var SubredditUser models.SubredditUser
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&SubredditUser)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newSubredditUser, error := m.DB.UpdateSubredditUser(value, SubredditUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersBySubredditUser(&newSubredditUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubredditUser)
}

func (m *Repository) DeleteSubredditUser(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	SubredditUser, error := m.DB.DeleteSubredditUser(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersBySubredditUser(&SubredditUser)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SubredditUser)
}

func getSubredditsAndUsersBySubredditUser(su *models.SubredditUser) error {
	var error error
	su.User, error = Repo.DB.GetUserById(su.UserId.String())
	if error != nil {
		return error
	}
	su.User.Password = "restricted"
	su.Subreddit, error = Repo.DB.GetSubredditById(su.SubredditId.String())
	if error != nil {
		return error
	}
	su.Subreddit.User, error = Repo.DB.GetUserById(su.Subreddit.CreatedBy.String())
	su.Subreddit.User.Password = "restricted"
	return error
}
