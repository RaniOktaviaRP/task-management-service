package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
    RoleSE  = "SE"
    RoleSCE = "SCE"
)

type User struct {
	Id           uuid.UUID  `json:"id"`
	FullName     string     `json:"full_name"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"password_hash"`
	Role         string     `json:"role"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}