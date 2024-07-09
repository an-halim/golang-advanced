package response

import "github.com/an-halim/golang-advanced/session-9-crud-user-grpc-gateway/entity"

type Success struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    entity.User `json:"data,omitempty"`
}
