package request

type Answers struct {
	QuestionID int    `json:"question_id" validate:"required,numeric"`
	Answer     string `json:"answer" validate:"required"`
}
