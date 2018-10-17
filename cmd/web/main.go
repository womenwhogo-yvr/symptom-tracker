package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define command line flags for configuration
	addr := flag.String("addr", ":3030", "HTTP network address")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static/", "Path to static assets")

	// Parse flags and assigns in the above variables
	flag.Parse()

	app := &App{
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
