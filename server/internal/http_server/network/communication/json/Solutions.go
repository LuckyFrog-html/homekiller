package json

type HomeworkSolveReviewJson struct {
	Comment string `json:"comment"`
	Score   int    `json:"score"`
}

type DeleteHomeworkSolveJson struct {
	HomeworkSolveId uint `json:"homework_solve_id"`
}
