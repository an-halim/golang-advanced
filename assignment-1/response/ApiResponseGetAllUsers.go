package response

type ApiResponseGetAllUsers struct {
	Page      int                      `json:"page"`
	Limit     int                      `json:"limit"`
	TotalPage int                      `json:"total_page"`
	Users     []ResponseSubmissionInfo `json:"users"`
}
