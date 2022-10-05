// Filename: cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// The application version number
const version = "1.0.0"

// create a configuration struct
type config struct {
	port int
	env  string // diffrent type of environment, development, production
}

// Dependency Injectiion, so its availabe to the handlers.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	//read in the flags that will be used to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production")
	flag.Parse() // need to do this step so we can access the flags
	//create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}
	//create our server
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//create our HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//we start our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
