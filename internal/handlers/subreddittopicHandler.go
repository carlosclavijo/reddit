package handlers

import (
	"encoding/json"
	"net/http"

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
