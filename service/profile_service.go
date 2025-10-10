package service

import (
	"context"
	"github.com/google/uuid"
	"task-management/model/web"
)

type ProfileService interface {
	Create(ctx context.Context, request web.ProfileCreateRequest) web.ProfileResponse
	Update(ctx context.Context, request web.ProfileUpdateRequest) web.ProfileResponse
	Delete(ctx context.Context, profileId uuid.UUID)
	FindById(ctx context.Context, profileId uuid.UUID) web.ProfileResponse
	FindByUserId(ctx context.Context, userId uuid.UUID) web.ProfileResponse
	FindAll(ctx context.Context) []web.ProfileResponse
}