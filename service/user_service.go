package service

import (
	"context"

	"github.com/google/uuid"
	"task-management/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error)

	// Login mengembalikan accessToken + refreshToken
	Login(ctx context.Context, request web.UserLoginRequest) (accessToken string, refreshToken string, err error)

	// Refresh token → generate access + refresh token baru
	Refresh(ctx context.Context, oldRefreshToken string) (newAccess string, newRefresh string, err error)

	// Logout → hapus refresh token
	Logout(ctx context.Context, refreshToken string) error

	Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error)
	Delete(ctx context.Context, userId uuid.UUID) error
	FindAll(ctx context.Context) ([]web.UserResponse, error)
	FindById(ctx context.Context, userId uuid.UUID) (web.UserResponse, error)
}
