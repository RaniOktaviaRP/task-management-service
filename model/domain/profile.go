package domain

import "github.com/google/uuid"

type Profile struct {
	Id       uuid.UUID
	UserId   uuid.UUID
	FullName string
	Email    string
	Role     string
}