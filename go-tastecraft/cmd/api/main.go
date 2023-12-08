package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"tastecraft/db/driver"
	"tastecraft/db/models"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		connectionString string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	DB       models.DBModel
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production|maintenance}")

	flag.Parse()

	connectionString := os.Getenv("CONNECTION_STRING")
	if len(connectionString) > 1 {
		connectionString = connectionString[1 : len(connectionString)-2]
	}

	cfg.db.connectionString = connectionString
	conn, err := driver.OpenDB(cfg.db.connectionString)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		DB:       models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting Back end server in %s on port %d", app.config.env, app.config.port)
	return srv.ListenAndServe()
}
