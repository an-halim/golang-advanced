package response

import (
	"time"
)

type GetOne struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	RiskScore      int       `json:"risk_score"`
	RiskCategory   string    `json:"risk_category"`
	RiskDefinition string    `json:"risk_definition"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
