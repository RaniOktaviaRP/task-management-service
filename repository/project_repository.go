package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"task-management/model/domain"
)

type ProjectRepository interface {
	Save(ctx context.Context, tx *sql.Tx, project domain.Project) domain.Project
	Update(ctx context.Context, tx *sql.Tx, project domain.Project) domain.Project
	Delete(ctx context.Context, tx *sql.Tx, projectId uuid.UUID) error
	FindById(ctx context.Context, tx *sql.Tx, projectId uuid.UUID) (domain.Project, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId uuid.UUID) []domain.Project
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Project
}
