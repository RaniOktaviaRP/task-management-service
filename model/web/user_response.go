package web

import "github.com/google/uuid"


type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	FullName string `json:"full_name"`
}
