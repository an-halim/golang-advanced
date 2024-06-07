package response

type APIResponseFailed struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
