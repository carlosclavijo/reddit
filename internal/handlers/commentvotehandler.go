package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostCommentVote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var CommentVote models.CommentVote
	err := decoder.Decode(&CommentVote)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newCommentVote, error := m.DB.InsertCommentVote(CommentVote)
	if error != nil {
		log.Println(error)
		helpers.ServerError(w, error)
		return
	}
	m.App.Session.Put(r.Context(), "commentvote", CommentVote)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCommentVote)
}
