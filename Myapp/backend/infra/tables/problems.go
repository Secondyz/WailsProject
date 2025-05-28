package tables

import "database/sql"

func CreateProblemsTable(db *sql.DB) error {
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS problems (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT
    );
    `

	_, err := db.Exec(sqlStmt)
	return err
}
