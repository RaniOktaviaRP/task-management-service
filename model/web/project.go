package web

import (
	"github.com/google/uuid"
	"time"
)

type ProjectCreateRequest struct {
	Name        string  `validate:"required" json:"name"`
	Description string  `json:"description"`
	Progress    float64 `validate:"min=0,max=100" json:"progress"`
	Confidence  float64 `validate:"min=0,max=100" json:"confidence"`
	Trend       string  `validate:"oneof=up down stable" json:"trend"`
	UserId      uuid.UUID `validate:"required" json:"user_id"`
}

type ProjectUpdateRequest struct {
	Id          uuid.UUID `validate:"required" json:"id"`
	Name        string    `validate:"required" json:"name"`
	Description string    `json:"description"`
	Progress    float64   `validate:"min=0,max=100" json:"progress"`
	Confidence  float64   `validate:"min=0,max=100" json:"confidence"`
	Trend       string    `validate:"oneof=up down stable" json:"trend"`
}

type ProjectResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Progress    float64   `json:"progress"`
	Confidence  float64   `json:"confidence"`
	Trend       string    `json:"trend"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      uuid.UUID `json:"user_id"`
}