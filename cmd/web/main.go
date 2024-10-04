package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/carlosclavijo/reddit/internal/config"
	"github.com/carlosclavijo/reddit/internal/driver"
	"github.com/carlosclavijo/reddit/internal/handlers"
	"github.com/carlosclavijo/reddit/internal/models"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func run() (*driver.DB, error) {
	gob.Register(models.User{})
	gob.Register(models.Topic{})
	gob.Register(models.Subreddit{})
	gob.Register(models.SubredditUser{})
	gob.Register(models.SubredditTopic{})
	gob.Register(models.Config{})
	gob.Register(models.Tag{})
	gob.Register(models.Topic{})
	gob.Register(models.Post{})
	gob.Register(models.PostTag{})
	gob.Register(models.Image{})
	gob.Register(models.Video{})
	gob.Register(models.Link{})
	gob.Register(models.Poll{})
	gob.Register(models.Option{})
	gob.Register(models.OptionUser{})
	gob.Register(models.Comment{})
	gob.Register(models.CommentVote{})

	app.Inproduction = false
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction
	app.Session = session

	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=reddit user=postgres password=abc12345")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	app.UseCache = false
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	return db, nil
}
