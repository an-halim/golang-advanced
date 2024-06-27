package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	ID           int             `json:"id"`
	UserID       int             `json:"user_id" gorm:"column:user_id"`
	RiskScore    int             `json:"risk_score"`
	RiskCategory string          `json:"risk_category"`
	Answers      json.RawMessage `json:"answers" gorm:"type:jsonb"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	User         User
}
