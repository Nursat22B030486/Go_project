package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/Nursat22B030486/go_project/pkg/read-it/model"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8888", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres@localhost:5432/readit", "PostgreSQL DSN")
	flag.Parse()

	// connect to DB

	db, err := openDB(cfg)

	if err3 := db.Ping(); err3 != nil {
		fmt.Println(err3)
		return
	}

	if err != nil {
		log.Fatal(err)
		fmt.Printf("NOONOOOOO")
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModule(db),
	}

	app.run()

}
func (app *application) run() {
	r := mux.NewRouter()

	// v1 := r.PathPrefix("/api/v1").Subrouter()

	r.HandleFunc("/articles", app.createArticleHandler).Methods("POST")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.getArticleHandler).Methods("GET")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.updateArticleHandler).Methods("PUT")
	r.HandleFunc("/articles/{articleId:[0-9]+}", app.deleteArticleHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err1 := http.ListenAndServe(app.config.port, r)
	log.Fatal(err1)

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=postgres password=pa55word dbname=readit sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
