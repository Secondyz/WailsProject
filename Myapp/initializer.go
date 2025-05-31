package main

import (
	"Myapp/api/problems"
	"Myapp/backend/app/judgetool/handler"
	"Myapp/backend/app/judgetool/service"
	"database/sql"
)

func InitializeHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		Problems: problems.NewProblemsAPI(
			handler.NewProblemsHandler(
				service.NewProblemService(db),
			),
		),
	}
}
