package domain

import (
    "time"
    "github.com/google/uuid"
)

type RefreshToken struct {
    Id        uuid.UUID
    UserID    uuid.UUID
    Token     string
    ExpiresAt time.Time
}

func NewRefreshToken(userID uuid.UUID, tokenString string) RefreshToken {
    return RefreshToken{
        Id:        uuid.New(),
        UserID:    userID,
        Token:     tokenString,
        ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 hari
    }
}
