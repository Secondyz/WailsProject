package problems

import (
	"Myapp/api/types"
	"Myapp/backend/app/judgetool/handler"
	"fmt"
)

type ProblemsAPI struct {
	problemsHandler *handler.ProblemsHandler
}

func NewProblemsAPI(problemsHandler *handler.ProblemsHandler) *ProblemsAPI {
	return &ProblemsAPI{
		problemsHandler: problemsHandler,
	}
}

func (api *ProblemsAPI) CreateProblem(req CreateProblemRequest) types.APIResponse[ProblemResponse] {
	if req.Title == "" {
		return types.APIResponse[ProblemResponse]{
			Success: false,
			Message: "タイトルを入力してください",
		}
	}

	err := api.problemsHandler.CreateProblem(handler.ProblemInput{
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		return types.APIResponse[ProblemResponse]{
			Success: false,
			Message: fmt.Sprintf("問題の作成に失敗しました: %v", err),
		}
	}

	return types.APIResponse[ProblemResponse]{
		Success: true,
		Message: "問題作成しました",
	}
}
