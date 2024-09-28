package handlers

import (
	"fmt"
	"net/http"

	"github.com/carlosclavijo/reddit/internal/config"
	"github.com/carlosclavijo/reddit/internal/driver"
	"github.com/carlosclavijo/reddit/internal/repository"
	"github.com/carlosclavijo/reddit/internal/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
