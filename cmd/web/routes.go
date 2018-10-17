package main

import (
	"net/http"
)

func (app *App) Routes() *http.ServeMux {
		// Use the http.NewServeMux() function to initialize a new serve mux (router)
		mux := http.NewServeMux()

		// Use the mux.HandleFunc() method to register function and handle URL pattern
		mux.HandleFunc("/", app.Home)
		mux.HandleFunc("/symptoms", app.ShowSymptoms)
	
		// create fileserver for static files
		// filepath should be relative to project directory
		fileServer := http.FileServer(http.Dir(app.StaticDir))

		mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	return mux
}