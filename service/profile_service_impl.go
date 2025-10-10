package service

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"task-management/helper"
	"task-management/model/domain"
	"task-management/model/web"
	"task-management/repository"
)

type ProfileServiceImpl struct {
	ProfileRepository repository.ProfileRepository
	DB               *sql.DB
}

func NewProfileService(profileRepository repository.ProfileRepository, db *sql.DB) ProfileService {
	return &ProfileServiceImpl{
		ProfileRepository: profileRepository,
		DB:               db,
	}
}

func (s *ProfileServiceImpl) Create(ctx context.Context, request web.ProfileCreateRequest) web.ProfileResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profile := domain.Profile{
		UserId:   request.UserId,
		FullName: request.FullName,
		Email:    request.Email,
		Role:     request.Role,
	}

	profile = s.ProfileRepository.Save(ctx, tx, profile)

	return toProfileResponse(profile)
}

func (s *ProfileServiceImpl) Update(ctx context.Context, request web.ProfileUpdateRequest) web.ProfileResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profile, err := s.ProfileRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	profile.UserId = request.UserId
	profile.FullName = request.FullName
	profile.Email = request.Email
	profile.Role = request.Role

	profile = s.ProfileRepository.Update(ctx, tx, profile)

	return toProfileResponse(profile)
}

func (s *ProfileServiceImpl) Delete(ctx context.Context, profileId uuid.UUID) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.ProfileRepository.FindById(ctx, tx, profileId)
	helper.PanicIfError(err)

	s.ProfileRepository.Delete(ctx, tx, profileId)
}

func (s *ProfileServiceImpl) FindById(ctx context.Context, profileId uuid.UUID) web.ProfileResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profile, err := s.ProfileRepository.FindById(ctx, tx, profileId)
	helper.PanicIfError(err)

	return toProfileResponse(profile)
}

func (s *ProfileServiceImpl) FindByUserId(ctx context.Context, userId uuid.UUID) web.ProfileResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profile, err := s.ProfileRepository.FindByUserId(ctx, tx, userId)
	helper.PanicIfError(err)

	return toProfileResponse(profile)
}

func (s *ProfileServiceImpl) FindAll(ctx context.Context) []web.ProfileResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	profiles := s.ProfileRepository.FindAll(ctx, tx)

	var profileResponses []web.ProfileResponse
	for _, profile := range profiles {
		profileResponses = append(profileResponses, toProfileResponse(profile))
	}

	return profileResponses
}

func toProfileResponse(profile domain.Profile) web.ProfileResponse {
	return web.ProfileResponse{
		Id:       profile.Id,
		UserId:   profile.UserId,
		FullName: profile.FullName,
		Email:    profile.Email,
		Role:     profile.Role,
	}
}