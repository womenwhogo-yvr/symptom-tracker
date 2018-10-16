package main

import (
	"net/http"
)

// Home : Handler for "/" route
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	// Use r.URL.Path to check whether the request URL path exactly matches "/".
	// If it doesn't, use the http.NotFound() function to send a 404 Not Found response.
	// Importantly, we then return from the function so that the subsequent code is not executed.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Symptom Tracker"))
}

// ShowSymptoms : Handler for "/symptoms" route
func (app *App) ShowSymptoms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of symptoms"))
}
