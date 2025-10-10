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

type ProjectServiceImpl struct {
	ProjectRepository repository.ProjectRepository
	DB                *sql.DB
}

func NewProjectService(projectRepository repository.ProjectRepository, db *sql.DB) ProjectService {
	return &ProjectServiceImpl{
		ProjectRepository: projectRepository,
		DB:                db,
	}
}

func (s *ProjectServiceImpl) Create(ctx context.Context, request web.ProjectCreateRequest) web.ProjectResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	project := domain.Project{
		Name:        request.Name,
		Description: request.Description,
		Progress:    request.Progress,
		Confidence:  request.Confidence,
		Trend:       request.Trend,
		UserId:      request.UserId,
	}

	project = s.ProjectRepository.Save(ctx, tx, project)

	return toProjectResponse(project)
}

func (s *ProjectServiceImpl) Update(ctx context.Context, request web.ProjectUpdateRequest) web.ProjectResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	project, err := s.ProjectRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	project.Name = request.Name
	project.Description = request.Description
	project.Progress = request.Progress
	project.Confidence = request.Confidence
	project.Trend = request.Trend

	project = s.ProjectRepository.Update(ctx, tx, project)

	return toProjectResponse(project)
}

func (s *ProjectServiceImpl) Delete(ctx context.Context, projectId uuid.UUID) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = s.ProjectRepository.FindById(ctx, tx, projectId)
	helper.PanicIfError(err)

	err = s.ProjectRepository.Delete(ctx, tx, projectId)
	helper.PanicIfError(err)
}

func (s *ProjectServiceImpl) FindById(ctx context.Context, projectId uuid.UUID) web.ProjectResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	project, err := s.ProjectRepository.FindById(ctx, tx, projectId)
	helper.PanicIfError(err)

	return toProjectResponse(project)
}

func (s *ProjectServiceImpl) FindByUserId(ctx context.Context, userId uuid.UUID) []web.ProjectResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	projects := s.ProjectRepository.FindByUserId(ctx, tx, userId)

	var projectResponses []web.ProjectResponse
	for _, project := range projects {
		projectResponses = append(projectResponses, toProjectResponse(project))
	}

	return projectResponses
}

func (s *ProjectServiceImpl) FindAll(ctx context.Context) []web.ProjectResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	projects := s.ProjectRepository.FindAll(ctx, tx)

	var projectResponses []web.ProjectResponse
	for _, project := range projects {
		projectResponses = append(projectResponses, toProjectResponse(project))
	}

	return projectResponses
}

func toProjectResponse(project domain.Project) web.ProjectResponse {
	return web.ProjectResponse{
		Id:          project.Id,
		Name:        project.Name,
		Description: project.Description,
		Progress:    project.Progress,
		Confidence:  project.Confidence,
		Trend:       project.Trend,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		UserId:      project.UserId,
	}
}
