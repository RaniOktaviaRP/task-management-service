package service

import (
	"context"
	"task-management/model/web"

	"github.com/google/uuid"
)

type TaskService interface {
	Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponse
	Update(ctx context.Context, taskId uuid.UUID, request web.TaskUpdateRequest) web.TaskResponse
	Delete(ctx context.Context, taskId uuid.UUID)
	FindById(ctx context.Context, taskId uuid.UUID) web.TaskResponse
	FindByProjectId(ctx context.Context, projectId uuid.UUID) []web.TaskResponse
	FindAll(ctx context.Context) []web.TaskResponse
}