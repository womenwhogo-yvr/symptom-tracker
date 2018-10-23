package main

import (
	"symptom-tracker/pkg/models"
)

// Define App struct to hold app wide dependencies and config settings
type App struct {
	Database *models.Database
	HTMLDir string
	StaticDir string
}