package web

import "github.com/google/uuid"


type UserUpdateRequest struct {
	Id       uuid.UUID `json:"id" validate:"required"`
	Email    *string   `json:"email"`            
	Password *string   `json:"password"`        
	Role     *string   `json:"role" validate:"omitempty,oneof=SE SCE"` 
	FullName *string   `json:"full_name"`
}

