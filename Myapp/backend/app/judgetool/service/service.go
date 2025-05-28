package service

import (
	"Myapp/backend/app/judgetool/models"
	"database/sql"
	"log"
)

type ProblemsService struct {
	DB *sql.DB
}

func NewProblemService(db *sql.DB) *ProblemsService {
	return &ProblemsService{DB: db}
}
func (s *ProblemsService) CreateProblem(p *models.Problem) error {
	stmt, err := s.DB.Prepare("INSERT INTO problems(title, description) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	_, err = stmt.Exec(p.Title, p.Description)
	return err
}
