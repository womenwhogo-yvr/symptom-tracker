package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"symptom-tracker/pkg/models"

	_ "github.com/lib/pq"
)

func main() {
	// Define command line flags for configuration
	addr := flag.String("addr", ":3030", "HTTP network address")
	dsn := flag.String("dsn", "dbname=symptom-tracker sslmode=disable", "Postgres data source name")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static/", "Path to static assets")

	// Parse flags and assigns in the above variables
	flag.Parse()

	// Open a connection to the db
	db := connect(*dsn)

	// Defer call ensures connection pool closes before the main() function exits
	defer db.Close()

	app := &App{
		Database: &models.Database{db},
		HTMLDir: *htmlDir,
		StaticDir: *staticDir,
	}

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":3030")
	// and the serve mux we just created. If ListenAndServe() returns an error we
	// use the log.Fatal() function to record the error message and exit the program.
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())

	log.Fatal(err)
}

// connect wraps sql.Open and returns a sql.DB connection pool
func connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}