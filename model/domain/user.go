package domain

import "github.com/google/uuid"

const (
    RoleSE  = "SE"
    RoleSCE = "SCE"
)

type User struct {
	Id           uuid.UUID
	FullName     string
	Email        string
	PasswordHash string
	Role         string
}