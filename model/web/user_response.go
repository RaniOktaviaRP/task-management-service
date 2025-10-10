package web

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	Id        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	FullName  string     `json:"full_name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
