package db

// DeleteSuperhero saves newly registered Superhero.
func(db *DB) DeleteSuperhero(id string) error {
	_, err := db.stmtDeleteSuperhero.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
