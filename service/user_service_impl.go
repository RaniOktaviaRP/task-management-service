package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"task-management/helper"
	"task-management/model/domain"
	"task-management/model/web"
	"task-management/repository"
)

type UserServiceImpl struct {
	UserRepository         repository.UserRepository
	ProfileRepository      repository.ProfileRepository
	RefreshTokenRepository repository.RefreshTokenRepository
	DB                     *sql.DB
	JwtSecret              []byte
}

// Constructor
func NewUserService(
	userRepository repository.UserRepository,
	profileRepository repository.ProfileRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	db *sql.DB,
	jwtSecret []byte,
) UserService {
	if db == nil {
		panic("DB cannot be nil")
	}
	return &UserServiceImpl{
		UserRepository:         userRepository,
		ProfileRepository:      profileRepository,
		RefreshTokenRepository: refreshTokenRepo,
		DB:                     db,
		JwtSecret:              jwtSecret,
	}
}

// Register user baru
func (s *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) (web.UserResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	email := strings.ToLower(strings.TrimSpace(request.Email))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserResponse{}, err
	}

	user := domain.User{
		Id:           uuid.New(),
		FullName:     strings.TrimSpace(request.FullName),
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         request.Role,
	}

	savedUser := s.UserRepository.Save(ctx, tx, user)

	// Buat profile untuk user baru
	profile := domain.Profile{
		UserId:   savedUser.Id,
		FullName: savedUser.FullName,
		Email:    savedUser.Email,
		Role:     savedUser.Role,
	}

	// Simpan profile
	s.ProfileRepository.Save(ctx, tx, profile)

	return web.UserResponse{
		Id:        savedUser.Id,
		FullName:  savedUser.FullName,
		Email:     savedUser.Email,
		Role:      savedUser.Role,
		CreatedAt: savedUser.CreatedAt,
		UpdatedAt: savedUser.UpdatedAt,
		DeletedAt: savedUser.DeletedAt,
	}, nil
}

// Login user → menghasilkan access + refresh token
func (s *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (accessToken string, refreshToken string, err error) {
	email := strings.ToLower(strings.TrimSpace(request.Email))

	user, err := s.UserRepository.FindByEmail(ctx, nil, email)
	if err != nil {
		return "", "", err
	}
	if user.Id == uuid.Nil {
		return "", "", errors.New("email tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return "", "", errors.New("password salah")
	}

	// Generate access token (1 jam)
	accessClaims := jwt.MapClaims{
		"user_id":   user.Id.String(),
		"full_name": user.FullName,
		"email":     user.Email,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = accessJWT.SignedString(s.JwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token (7 hari)
	refreshClaims := jwt.MapClaims{
		"user_id": user.Id.String(),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = refreshJWT.SignedString(s.JwtSecret)
	if err != nil {
		return "", "", err
	}

	rt := domain.NewRefreshToken(user.Id, refreshToken)
	err = s.RefreshTokenRepository.Save(ctx, nil, rt)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Refresh → generate access + refresh token baru
func (s *UserServiceImpl) Refresh(ctx context.Context, oldRefreshToken string) (newAccess string, newRefresh string, err error) {
	tokenData, err := s.RefreshTokenRepository.FindByToken(ctx, nil, oldRefreshToken)
	if err != nil {
		return "", "", err
	}

	user, err := s.UserRepository.FindById(ctx, nil, tokenData.UserID)
	if err != nil {
		return "", "", err
	}

	// Generate access token baru
	accessClaims := jwt.MapClaims{
		"user_id":   user.Id.String(),
		"full_name": user.FullName,
		"email":     user.Email,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	accessJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	newAccess, err = accessJWT.SignedString(s.JwtSecret)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token baru
	refreshClaims := jwt.MapClaims{
		"user_id": user.Id.String(),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	newRefresh, err = refreshJWT.SignedString(s.JwtSecret)
	if err != nil {
		return "", "", err
	}

	newRT := domain.NewRefreshToken(user.Id, newRefresh)
	_ = s.RefreshTokenRepository.Delete(ctx, nil, oldRefreshToken)
	_ = s.RefreshTokenRepository.Save(ctx, nil, newRT)

	return newAccess, newRefresh, nil
}

// Logout → hapus refresh token
func (s *UserServiceImpl) Logout(ctx context.Context, refreshToken string) error {
	return s.RefreshTokenRepository.Delete(ctx, nil, refreshToken)
}

// Update user
func (s *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) (web.UserResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	existingUser, err := s.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.UserResponse{}, err
	}

	if request.FullName != nil {
		existingUser.FullName = strings.TrimSpace(*request.FullName)
	}

	if request.Email != nil {
		existingUser.Email = strings.ToLower(strings.TrimSpace(*request.Email))
	}

	if request.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*request.Password), bcrypt.DefaultCost)
		if err != nil {
			return web.UserResponse{}, err
		}
		existingUser.PasswordHash = string(hashedPassword)
	}

	if request.Role != nil {
		existingUser.Role = *request.Role
	}

	updatedUser := s.UserRepository.Update(ctx, tx, existingUser)

	return web.UserResponse{
		Id:        updatedUser.Id,
		FullName:  updatedUser.FullName,
		Email:     updatedUser.Email,
		Role:      updatedUser.Role,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
		DeletedAt: updatedUser.DeletedAt,
	}, nil
}

// Delete user
func (s *UserServiceImpl) Delete(ctx context.Context, userId uuid.UUID) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	s.UserRepository.Delete(ctx, tx, userId)
	return nil
}

// FindById user
func (s *UserServiceImpl) FindById(ctx context.Context, userId uuid.UUID) (web.UserResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := s.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return web.UserResponse{}, err
	}

	return web.UserResponse{
		Id:        user.Id,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}

// FindAll users
func (s *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserResponse, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	users := s.UserRepository.FindAll(ctx, tx)
	responses := make([]web.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, web.UserResponse{
			Id:        user.Id,
			FullName:  user.FullName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		})
	}
	return responses, nil
}
