package tables

import "database/sql"

func CreateTestCasesTable(db *sql.DB) error {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS test_cases (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    problem_id INTEGER NOT NULL ,
	    input TEXT NOT NULL ,
	    expected_output TEXT NOT NULL ,
	    FOREIGN KEY (problem_id) REFERENCES problems(id)
	);
`
	_, err := db.Exec(sqlStmt)
	return err
}
