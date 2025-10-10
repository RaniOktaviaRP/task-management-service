package repository

import (
    "context"
    "database/sql"
    "errors"
    "task-management/model/domain"
    "time"
)

type RefreshTokenRepositoryImpl struct {
}

// Constructor
func NewRefreshTokenRepository() RefreshTokenRepository {
    return &RefreshTokenRepositoryImpl{}
}

// Simulasi map (ganti nanti ke query PostgreSQL)
var refreshTokens = map[string]domain.RefreshToken{}

func (r *RefreshTokenRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, token domain.RefreshToken) error {
    refreshTokens[token.Token] = token
    return nil
}

func (r *RefreshTokenRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, tokenString string) (domain.RefreshToken, error) {
    token, ok := refreshTokens[tokenString]
    if !ok {
        return domain.RefreshToken{}, errors.New("refresh token not found")
    }
    if token.ExpiresAt.Before(time.Now()) {
        delete(refreshTokens, tokenString)
        return domain.RefreshToken{}, errors.New("refresh token expired")
    }
    return token, nil
}

func (r *RefreshTokenRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, tokenString string) error {
    delete(refreshTokens, tokenString)
    return nil
}
