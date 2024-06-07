package response

import "github.com/an-halim/golang-advanced/session-2-latihan-crud/entity"

type APIResponseGetAll struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Page    int           `json:"page"`
	Limit   int           `json:"limit"`
	Data    []entity.User `json:"data"`
}
