package tables

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	sqlSTmt := `
	CREATE TABLE IF NOT EXISTS users(
	    id INTEGER PRIMARY KEY  AUTOINCREMENT ,
	    username TEXT NOT NULL UNIQUE ,
	    email TEXT
	    );
`
	_, err := db.Exec(sqlSTmt)
	return err
}
