package response

import "github.com/an-halim/golang-advanced/session-2-latihan-crud/entity"

type APIResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    entity.User `json:"data"`
}
