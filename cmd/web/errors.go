package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

// ServerError writes an error and stack trace for the current goroutine
// to the log and returns a server error response to the user
func (app *App) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// ClientError sends a specific error and status code to the user
// e.g. 400 Bad Request
func (app *App) ClientError(w http.ResponseWriter, status int) {
	// http.StatusText returns a human-friendly text representation of status code
	http.Error(w, http.StatusText(status), status)
}

// NotFound is a convenience wrapper around ClientError that sends a
// 404 response to the user
func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}