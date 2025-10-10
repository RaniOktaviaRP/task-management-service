package repository

import (
    "context"
    "database/sql"
    "task-management/model/domain"
)

type RefreshTokenRepository interface {
    Save(ctx context.Context, tx *sql.Tx, token domain.RefreshToken) error
    FindByToken(ctx context.Context, tx *sql.Tx, tokenString string) (domain.RefreshToken, error)
    Delete(ctx context.Context, tx *sql.Tx, tokenString string) error
}
