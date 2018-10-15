package main

import (
	"net/http"
)

// Home : Handler for "/" route
func Home(w http.ResponseWriter, r *http.Request) {
	// Use r.URL.Path to check whether the request URL path exactly matches "/".
	// If it doesn't, use the w.WriteHeader() method to send a 404 status code and
	// the w.Write() method to write a "Not Found" response body. Importantly, we
	// then return from the function so that the subsequent code is not executed.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Symptom Tracker"))
}

// ShowSymptoms : Handler for "/symptoms" route
func ShowSymptoms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of symptoms"))
}
