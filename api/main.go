package main

import (
	"FitnessTracker/internal/data"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	config config
	logger *log.Logger
	mux    *http.ServeMux
	models data.Models
}

type config struct {
	port int
	env  string
	dsn  string
}

const appVersion = "1.0.0"

func main() {

	var cfg config

	// setting up config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production")
	flag.StringVar(&cfg.dsn, "dsn", "sqlite:userdb", "dsn URI")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mux := http.NewServeMux()

	app := &application{
		config: cfg,
		logger: logger,
		mux:    mux,
	}

	conn := app.setupDB()
	app.models = data.NewModels(conn)

	setupUsersTable(conn)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// printing out the starting configuration
	logger.Printf("Starting server %s on port %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
