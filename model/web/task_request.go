package web

import (
"github.com/google/uuid"
"time"
)

type TaskCreateRequest struct {
	ProjectId      uuid.UUID `json:"project_id" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Status         string    `json:"status" validate:"omitempty,oneof=todo in-progress completed"`
	Priority       string    `json:"priority" validate:"omitempty,oneof=low medium high"`
	Effort         int       `json:"effort" validate:"required"`
	DifficultyLevel string    `json:"difficulty_level"`
	Deliverable    string    `json:"deliverable"`
	Bottleneck     string    `json:"bottleneck"`
}

type TaskUpdateRequest struct {
	Title          string    `json:"title" validate:"required"`
	Status         string    `json:"status" validate:"required,oneof=todo in-progress completed"`
	Priority       string    `json:"priority" validate:"required,oneof=low medium high"`
	Effort         int       `json:"effort" validate:"required"`
	DifficultyLevel string    `json:"difficulty_level"`
	Deliverable    string    `json:"deliverable"`
	Bottleneck     string    `json:"bottleneck"`
}

type TaskResponse struct {
	Id             uuid.UUID `json:"id"`
	ProjectId      uuid.UUID `json:"project_id"`
	Title          string    `json:"title"`
	Status         string    `json:"status"`
	Priority       string    `json:"priority"`
	Effort         int       `json:"effort"`
	DifficultyLevel string    `json:"difficulty_level"`
	Deliverable    string    `json:"deliverable"`
	Bottleneck     string    `json:"bottleneck"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}