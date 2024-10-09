package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetConfigsList(w http.ResponseWriter, r *http.Request) {
	configs, error := m.DB.GetConfigs()
	if error != nil {
		helpers.ServerError(w, error)
	}
	for i := 0; i < len(configs); i++ {
		error = getSubredditsAndUsersByConfig(&configs[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(configs)
}

func (m *Repository) GetConfigById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	config, error := m.DB.GetConfigById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByConfig(&config)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

func (m *Repository) PostConfig(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Config models.Config
	err := decoder.Decode(&Config)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newConfig, error := m.DB.InsertConfig(Config)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	newConfig.Admin.Password = "restricted"
	//m.App.Session.Put(r.Context(), "config", Config)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newConfig)
}

func (m *Repository) PutConfig(w http.ResponseWriter, r *http.Request) {
	var Config models.Config
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&Config)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newConfig, error := m.DB.UpdateConfig(value, Config)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByConfig(&newConfig)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newConfig)
}

func (m *Repository) DeleteConfig(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	Config, error := m.DB.DeleteConfig(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByConfig(&Config)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Config)
}

func getSubredditsAndUsersByConfig(c *models.Config) error {
	var error error
	c.Admin, error = Repo.DB.GetUserById(c.AdminId.String())
	if error != nil {
		return error
	}
	c.Admin.Password = "restricted"
	c.Subreddit, error = Repo.DB.GetSubredditById(c.SubredditId.String())
	if error != nil {
		return error
	}
	c.Subreddit.User, error = Repo.DB.GetUserById(c.Subreddit.CreatedBy.String())
	c.Subreddit.User.Password = "restricted"
	return error
}
