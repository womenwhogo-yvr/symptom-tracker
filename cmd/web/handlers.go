package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home : Handler for "/" route
func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	// Use r.URL.Path to check whether the request URL path exactly matches "/".
	// If it doesn't, use the http.NotFound() function to send a 404 Not Found response.
	// Importantly, we then return from the function so that the subsequent code is not executed.
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	// Combines HTML templates and takes any dynamic data if available
	app.RenderHTML(w)
}

// ShowEntry : Handler for "/entry" route
func (app *App) ShowEntry(w http.ResponseWriter, r *http.Request) {
	// get the query param for id and convert to integer
	// with strconv.Atoi
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	entry, err := app.Database.GetEntry(id)
	symptoms, err := app.Database.GetSymptoms(id)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	fmt.Fprint(w, entry)
	for _, p := range symptoms {
    fmt.Printf("%+v\n", p)
	}
}

// AddEntry : Handler for "/entry/add" route
func (app *App) AddEntry(w http.ResponseWriter, r *http.Request) {
	// TODO
}