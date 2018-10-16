package main

import (
	"html/template"
	"net/http"
	"log"
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

	// Initialize slice containing paths to template files
	files := []string{
		"./ui/html/layout.html",
	}

	// Read and store files in template set
	// If error, return response with standard message & status code
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// use ExecuteTemplate to use "layout" template and write content to
	// http.ResponseWriter
	// 3rd parameter represents any dynamic data to be passed
	err = ts.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// ShowSymptoms : Handler for "/symptoms" route
func (app *App) ShowSymptoms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of symptoms"))
}
