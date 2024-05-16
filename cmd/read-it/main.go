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

	"github.com/Nursat22B030486/go_project/pkg/read-it/jsonlog"
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
	logger *jsonlog.Logger
	models model.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8888, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgresql://postgres:pa55word@localhost:5432/readit?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Initialization of new looger
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

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

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		models: model.NewModule(db),
		logger: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		// Create a new Go log.Logger instance with the log.New() function, passing in
		// our custom Logger as the first parameter. The "" and 0 indicate that the
		// log.Logger instance should not use a prefix or any flags.
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Again, we use the PrintInfo() method to write a "starting server" message at the
	// INFO level. But this time we pass a map containing additional properties (the
	// operating environment and server address) as the final parameter.
	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  cfg.env,
	})

	err1 := srv.ListenAndServe()
	logger.PrintFatal(err1, nil)

}

func openDB(cfg config) (*sql.DB, error) {
	// db, err := sql.Open("postgres", "user=postgres password=pa55word dbname=readit sslmode=disable")
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
