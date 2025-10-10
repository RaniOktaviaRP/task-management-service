package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"task-management/model/domain"
)

type ProfileRepository interface {
	Save(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile
	Update(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile
	Delete(ctx context.Context, tx *sql.Tx, profileId uuid.UUID)
	FindById(ctx context.Context, tx *sql.Tx, profileId uuid.UUID) (domain.Profile, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, userId uuid.UUID) (domain.Profile, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Profile
}