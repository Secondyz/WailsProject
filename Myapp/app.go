package main

import (
	"Myapp/api/problems"
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
	Problems *problems.ProblemsAPI
}

// NewApp creates a new App application struct
func NewApp() *App {

	db, err := sql.Open("sqlite3", "./data/judge.db")
	if err != nil {
		panic(err)
	}
	handlers := InitializeHandlers(db)
	return &App{
		handlers: handlers,
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
