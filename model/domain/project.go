package domain

import (
	"github.com/google/uuid"
	"time"
)

type Project struct {
	Id          uuid.UUID
	Name        string
	Description string
	Progress    float64
	Confidence  float64
	Trend       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      uuid.UUID
}