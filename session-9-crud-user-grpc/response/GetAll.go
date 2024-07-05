package response

import "github.com/an-halim/golang-advanced/session-9-crud-user-grpc/entity"

type GetAll struct {
	Status      string        `json:"status"`
	Data        []entity.User `json:"data"`
	PageSize    int           `json:"page_size"`
	CurrentPage int           `json:"current_page"`
}
