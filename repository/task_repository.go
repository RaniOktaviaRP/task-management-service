package repository

import (
	"context"
	"task-management/model/domain"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Save(ctx context.Context, task domain.Task) (domain.Task, error)
	Update(ctx context.Context, task domain.Task) (domain.Task, error)
	Delete(ctx context.Context, taskId uuid.UUID) error
	FindById(ctx context.Context, taskId uuid.UUID) (domain.Task, error)
	FindByProjectId(ctx context.Context, projectId uuid.UUID) ([]domain.Task, error)
	FindAll(ctx context.Context) ([]domain.Task, error)
}