package main

import (
	"Myapp/backend/app/judgetool/handler"
	"Myapp/backend/app/judgetool/service"
	"Myapp/backend/infra/db"
	"context"
	"database/sql"
	"fmt"
)

// App struct
type App struct {
	ctx      context.Context
	handlers *Handlers
}
type Handlers struct {
	Problems *handler.ProblemsHandler
}

func setupHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		Problems: handler.NewProblemsHandler(service.NewProblemService(db)),
	}
}

// NewApp creates a new App application struct
func NewApp() *App {
	conn, err := db.OpenDB()
	if err != nil {
		//log.Fatal(err)
	}

	return &App{
		handlers: setupHandlers(conn),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type ProblemInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (a *App) CreateProblem(input ProblemInput) error {
	return a.handlers.Problems.CreateProblem(handler.ProblemInput{
		Title:       input.Title,
		Description: input.Description,
	})
}
