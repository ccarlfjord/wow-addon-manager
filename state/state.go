package state

import "database/sql"

// State is the primary interface for storing addon states in database
type State interface {
	Read()
	Write()
	Update()
	ReadAll()
}

func Read(s *State) {

}

func Write(s *State) error {
	return nil
}

func open(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
