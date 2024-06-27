package response

type GetAll struct {
	Status      string                   `json:"status"`
	Data        []ResponseSubmissionInfo `json:"data"`
	PageSize    int                      `json:"page_size"`
	CurrentPage int                      `json:"current_page"`
}
