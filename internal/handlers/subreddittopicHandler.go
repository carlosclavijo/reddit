package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) GetSubredditsTopicsList(w http.ResponseWriter, r *http.Request) {
	subredditsTopics, error := m.DB.GetSubredditsTopics()
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(subredditsTopics); i++ {
		error = getSubredditsAndTopicsBySubredditTopic(&subredditsTopics[i])
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subredditsTopics)
}

func (m *Repository) GetSubredditTopicById(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	subredditTopic, error := m.DB.GetSubredditsTopicById(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndTopicsBySubredditTopic(&subredditTopic)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subredditTopic)
}

func (m *Repository) GetSubredditsByTopic(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	subreddits, error := m.DB.GetSubredditsByTopicId(value)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	for i := 0; i < len(subreddits); i++ {
		subreddits[i].User, error = m.DB.GetUserById(subreddits[i].CreatedBy.String())
		if error != nil {
			helpers.ServerError(w, error)
			return
		}
		subreddits[i].User.Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subreddits)
}

func (m *Repository) GetTopicsBySubreddit(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[3]
	topics, error := m.DB.GetTopicsBySubredditId(value)
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
		topics[i].User.Password = "restricted"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topics)
}

func (m *Repository) PostSubedditTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var SubredditTopic models.SubredditTopic
	err := decoder.Decode(&SubredditTopic)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newSubredditTopic, error := m.DB.InsertSubredditTopic(SubredditTopic)
	if error != nil {
		helpers.ServerError(w, error)

	}
	error = getSubredditsAndTopicsBySubredditTopic(&newSubredditTopic)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "subreddittopic", SubredditTopic)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSubredditTopic)
}

func (m *Repository) DeleteSubredditTopic(w http.ResponseWriter, r *http.Request) {
	value := strings.Split(r.URL.Path, "/")[2]
	SubredditTopic, error := m.DB.DeleteSubredditTopic(value)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
		return
	}
	error = getSubredditsAndTopicsBySubredditTopic(&SubredditTopic)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	//m.App.Session.Put(r.Context(), "user", User)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SubredditTopic)
}

func getSubredditsAndTopicsBySubredditTopic(s *models.SubredditTopic) error {
	var error error
	s.Subreddit, error = Repo.DB.GetSubredditById(s.SubredditId.String())
	if error != nil {
		return error
	}
	s.Topic, error = Repo.DB.GetTopicById(s.TopicId.String())
	if error != nil {
		return error
	}
	s.Subreddit.User, error = Repo.DB.GetUserById(s.Subreddit.CreatedBy.String())
	if error != nil {
		return error
	}
	s.Topic.User, error = Repo.DB.GetUserById(s.Topic.UserId.String())
	return error
}
