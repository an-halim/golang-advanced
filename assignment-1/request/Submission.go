package request

type CreateSubmissionInfo struct {
	UserId  int       `json:"user_id" validate:"required,numeric"`
	Answers []Answers `json:"answers" validate:"required,dive"`
}
