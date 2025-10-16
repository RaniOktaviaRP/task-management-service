package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProjectId      uuid.UUID `gorm:"type:uuid;not null"`
	Title          string    `gorm:"type:text;not null"`
	Status         string    `gorm:"type:text;default:'todo'"`
	Priority       string    `gorm:"type:text;default:'medium'"`
	Effort         int       `gorm:"not null"`
	DifficultyLevel string    `gorm:"type:text"`
	Deliverable    string    `gorm:"type:text"`
	Bottleneck     string    `gorm:"type:text"`
	Progress       string    `gorm:"type:text"`
	ContinueTomorrow bool     `gorm:"type:boolean;default:false"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}