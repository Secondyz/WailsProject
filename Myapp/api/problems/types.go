package problems

type CreateProblemRequest struct {
	Title       string `json:"title" validate:"required, min=1, max=100"`
	Description string `json:"description" validate:"required, min=1"`
}

type ProblemResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TimeLimit   int    `json:"time_limit"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
