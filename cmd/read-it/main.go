package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/Nursat22B030486/go_project/pkg/read-it/model"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8888, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgresql://postgres:pa55word@localhost:5432/readit?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Initialization of new looger 
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)


	// connect to DB
	db, err := openDB(cfg)

	// if err3 := db.Ping(); err3 != nil {
	// 	fmt.Println(err3)
	// 	return
	// }

	if err != nil {
		log.Fatal(err)
		fmt.Printf("NOONOOOOO")
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModule(db),
		logger: logger,
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s\n", cfg.env, srv.Addr)
	err1 := srv.ListenAndServe()
	logger.Fatal(err1)

}


func openDB(cfg config) (*sql.DB, error) {
	// db, err := sql.Open("postgres", "user=postgres password=pa55word dbname=readit sslmode=disable")
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
