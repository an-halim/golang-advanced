package entity

import "time"

type User struct {
	ID          int          `json:"id"`
	Name        string       `json:"name" binding:"required"`
	Email       string       `json:"email" binding:"required,email"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Submissions []Submission `json:"submissions" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
