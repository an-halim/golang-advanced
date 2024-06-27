package response

import (
	"encoding/json"
	"time"
)

type GetOneSubmission struct {
	ID             int             `json:"id"`
	UserID         int             `json:"user_id"`
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	RiskScore      int             `json:"risk_score"`
	RiskCategory   string          `json:"risk_category"`
	RiskDefinition string          `json:"risk_definition"`
	Answer         json.RawMessage `json:"answer"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
