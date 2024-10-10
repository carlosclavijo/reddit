package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetTagsList(w http.ResponseWriter, r *http.Request) {
	tags, error := m.DB.GetTags()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(tags); i++ {
		error = getSubredditsAndUsersByTag(&tags[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (m *Repository) GetTagById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	tag, error := m.DB.GetTagById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByTag(&tag)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tag)
}

func (m *Repository) GetTagsBySubreddit(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	tags, error := m.DB.GetTagsBySubredditId(value)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(tags); i++ {
		error = getSubredditsAndUsersByTag(&tags[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

func (m *Repository) PostTag(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Tag models.Tag
	err := decoder.Decode(&Tag)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newTag, error := m.DB.InsertTag(Tag)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByTag(&newTag)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "tag", Tag)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTag)
}

func (m *Repository) PutTag(w http.ResponseWriter, r *http.Request) {
	var Tag models.Tag
	value := strings.Split(r.URL.Path, "/")[2]
	err := json.NewDecoder(r.Body).Decode(&Tag)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newTag, error := m.DB.UpdateTag(value, Tag)
	if error != nil {
		log.Print(error)
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTag)
}

func (m *Repository) DeleteTag(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	Tag, error := m.DB.DeleteTag(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndUsersByTag(&Tag)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tag)
}

func getSubredditsAndUsersByTag(t *models.Tag) error {
	var error error
	t.Admin, error = Repo.DB.GetUserById(t.AdminId.String())
	if error != nil {
		return error
	}
	t.Admin.Password = "restricted"
	t.Subreddit, error = Repo.DB.GetSubredditById(t.SubredditId.String())
	if error != nil {
		return error
	}
	t.Subreddit.User, error = Repo.DB.GetUserById(t.Subreddit.CreatedBy.String())
	t.Subreddit.User.Password = "restricted"
	return error
}
