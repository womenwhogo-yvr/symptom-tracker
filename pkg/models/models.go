package models

import (
	"time"
)

type User struct {
	ID					int
	UUID				string
	Name				string
	Email				string
	Password		string
	CreatedAt		time.Time
}

// type Session struct {
// 	ID					int
// 	UUID				string
// 	UserId			int
// 	Email				string
// 	CreatedAt		time.Time
// }

// Symptom: user added symptom which corresponds to an entry
type Symptom struct {
	ID						int
	Symptom				string
	Intensity			int
	EntryId				int
}

// Symptoms: slice that holds multiple Symptom type objects
type Symptoms []*Symptom

// Entry: user entry for a single date
type Entry struct {
	ID					int
  UserId			int
	Notes				string
	EntryDate		time.Time
}