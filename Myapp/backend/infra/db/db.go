package db

import (
	"Myapp/backend/infra/tables"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB() {
	db, err := OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err := tables.CreateProblemsTable(db); err != nil {
		log.Fatalf("failed to create problems table: %v", err)
	}
	if err := tables.CreateUsersTable(db); err != nil {
		log.Fatalf("failed to create users table: %v", err)
	}
	if err := tables.CreateSubmissionsTable(db); err != nil {
		log.Fatalf("failed to create submissions table: %v", err)
	}
	if err := tables.CreateTestCasesTable(db); err != nil {
		log.Fatalf("failed to create test_cases table: %v", err)
	}

}

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./data/judge.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
