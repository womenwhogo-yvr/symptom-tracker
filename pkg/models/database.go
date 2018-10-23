package models

import (
	"database/sql"
)

type Database struct {
	*sql.DB // embed connection pool for methods to gain access
}

// GetEntry takes an EntryId and queries the entries table
// for an entry that matches the id.
func (db *Database) GetEntry(id int) (*Entry, error) {
	statement := `SELECT id, entry_date, notes FROM entries
	WHERE id = $1`

	row := db.QueryRow(statement, id)

	entry := &Entry{}

	err := row.Scan(&entry.ID, &entry.EntryDate, &entry.Notes)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return entry, nil
}

// GetSymptoms takes an EntryId as a param and queries the 
// symptoms table for all user symptoms that match the EntryId.
func (db *Database) GetSymptoms(id int) (Symptoms, error) {
	statement := `SELECT id, symptom, intensity FROM symptoms
	WHERE entry_id = $1`

	rows, err := db.Query(statement, id)
	if err != nil {
		return nil, err
	}

	// Need to ensure that the connection is closed when GetSymptoms()
	// has completed. This should always come after checking for err.
	defer rows.Close()

	symptoms := Symptoms{}

	for rows.Next() {
		// Need to create pointer because Symptoms expects pointer to struct
		s := &Symptom{}

		err := rows.Scan(&s.ID, &s.Symptom, &s.Intensity)
		if err != nil {
			return nil, err
		}

		symptoms = append(symptoms, s)
	}

	// Checks once rows.Next() is complete to ensure no errors were
	// encountered. Cannot assume successful iteration over whole dataset.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return symptoms, nil
}

// InsertSymptoms takes user input and inserts into the symptoms table
func (db *Database) InsertSymptom(symptom string, intensity int, id int) error {
	statement := "INSERT INTO symptoms (symptom, intensity, entry_id) VALUES $1, $2, $3"

	// db.Exec inserts the values following statement into the symptoms table
	_, err := db.Exec(statement, symptom, intensity, id)
	if err != nil {
		return err
	}

	return nil
}