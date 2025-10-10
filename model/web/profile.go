package web

import "github.com/google/uuid"

type ProfileCreateRequest struct {
	UserId   uuid.UUID `json:"user_id" validate:"required"`
	FullName string    `json:"full_name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Role     string    `json:"role" validate:"required,oneof=SE SCE"`
}

type ProfileUpdateRequest struct {
	Id       uuid.UUID `json:"id" validate:"required"`
	UserId   uuid.UUID `json:"user_id" validate:"required"`
	FullName string    `json:"full_name" validate:"required"`
	Email    string    `json:"email" validate:"required,email"`
	Role     string    `json:"role" validate:"required,oneof=SE SCE"`
}

type ProfileResponse struct {
	Id       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"user_id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}