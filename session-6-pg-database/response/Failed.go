package response

type Failed struct {
	Status  string `json:"status"`
	Message string `json:"error"`
}
