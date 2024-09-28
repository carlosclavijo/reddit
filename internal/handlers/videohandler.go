package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/helpers"
	"github.com/carlosclavijo/reddit/internal/models"
)

func (m *Repository) PostVideo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Video models.Video
	err := decoder.Decode(&Video)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	newVideo, error := m.DB.InsertVideo(Video)
	if error != nil {
		helpers.ServerError(w, error)
		return
	}
	m.App.Session.Put(r.Context(), "video", Video)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newVideo)
}
