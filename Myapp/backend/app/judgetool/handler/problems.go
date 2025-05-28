package handler

import (
	"Myapp/backend/app/judgetool/models"
	"Myapp/backend/app/judgetool/service"
)

type ProblemsHandler struct {
	Service *service.ProblemsService
}

func NewProblemsHandler(service *service.ProblemsService) *ProblemsHandler {
	return &ProblemsHandler{Service: service}
}

type ProblemInput struct {
	Title       string
	Description string
}

func (h *ProblemsHandler) CreateProblem(input ProblemInput) error {
	problem := models.Problem{
		Title:       input.Title,
		Description: input.Description,
	}
	return h.Service.CreateProblem(&problem)
}
