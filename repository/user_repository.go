package repository

import (
	"context"
	"database/sql"
		"github.com/google/uuid"
	"task-management/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId uuid.UUID)
	FindById(ctx context.Context, tx *sql.Tx, userId uuid.UUID) (domain.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
