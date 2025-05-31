package main

import (
	"Myapp/api/problems"
	"Myapp/api/types"
)

func (a *App) CreateProblem(req problems.CreateProblemRequest) types.APIResponse[problems.ProblemResponse] {
	return a.handlers.Problems.CreateProblem(req)
}
