package tables

import "database/sql"

func CreateSubmissionsTable(db *sql.DB) error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS submissions(
	    id INTEGER PRIMARY KEY AUTOINCREMENT ,
	    user_id INTEGER NOT NULL ,
	    problem_id INTEGER NOT NULL,
	    source_code TEXT NOT NULL ,
	    language TEXT NOT NULL ,
	    submit_time DATETIME DEFAULT CURRENT_TIMESTAMP,
	    result TEXT,
	    FOREIGN KEY (user_id) REFERENCES users(id),
	    FOREIGN KEY (problem_id) REFERENCES problems(id)
	);
`
	_, err := db.Exec(sqlStmt)
	return err
}
