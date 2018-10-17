package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderHTML combines the template files with dynamic data and
// writes to response
func (app *App) RenderHTML(w http.ResponseWriter) {
	// Initialize slice containing paths to template files
	files := []string{
		filepath.Join(app.HTMLDir, "layout.html"),
		// TODO: Update when templates are available add additional arguments
		// filepath.Join(app.HTMLDir, page),
	}

	// Read and store files in template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	// Use ExecuteTemplate to use "layout" template and write content to
	// http.ResponseWriter
	// 3rd parameter represents any dynamic data to be passed
	err = ts.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		app.ServerError(w, err)
	}
}