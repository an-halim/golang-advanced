package response

import "github.com/an-halim/golang-advanced/assignment-1/entity"

type ApiResponseGetAllSubmission struct {
	UserID          int                 `json:"user_id,omitempty"`
	Page            int                 `json:"page"`
	Limit           int                 `json:"limit"`
	TotalSubmission int                 `json:"total_submissions,omitempty"`
	TotalPage       int                 `json:"total_page"`
	Submission      []entity.Submission `json:"submissions"`
}
