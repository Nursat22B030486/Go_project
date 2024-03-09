package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":1111", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "jdbc:postgresql://localhost:5432/postgres", "PostgreSQL DSN")
	flag.Parse()

	// connect to DB

	db, err := openDB(cfg)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/articles", CreateArticleHandler).Methods("POST")
	v1.HandleFunc("/articles/{articleId:[0-9]+}", GetArticleHandler).Methods("Get")
	v1.HandleFunc("/articles/{articleId:[0-9]+}", UpdateArticleHandler).Methods("PUT")
	v1.HandleFunc("/articles/{articleId:[0-9]+}", DeleteArticleHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", cfg.port)
	err1 := http.ListenAndServe(cfg.port, r)
	log.Fatal(err1)

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
