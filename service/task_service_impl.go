package service

import (
	"context"
	"time"

	"task-management/helper"
	"task-management/model/domain"
	"task-management/model/web"
	"task-management/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	Validator      *validator.Validate
}

func NewTaskService(taskRepository repository.TaskRepository, validator *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		Validator:      validator,
	}
}

func (service *TaskServiceImpl) Create(ctx context.Context, request web.TaskCreateRequest) web.TaskResponse {
	// Validasi request
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	// Generate UUID baru biar gak duplicate
	newID := uuid.New()

	task := domain.Task{
		Id:              newID,
		ProjectId:       request.ProjectId,
		Title:           request.Title,
		Status:          request.Status,
		Priority:        request.Priority,
		Effort:          request.Effort,
		DifficultyLevel: request.DifficultyLevel,
		Deliverable:     request.Deliverable,
		Bottleneck:      request.Bottleneck,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	result, err := service.TaskRepository.Save(ctx, task)
	helper.PanicIfError(err)

	return helper.ToTaskResponse(result)
}

func (service *TaskServiceImpl) Update(ctx context.Context, taskId uuid.UUID, request web.TaskUpdateRequest) web.TaskResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	task, err := service.TaskRepository.FindById(ctx, taskId)
	helper.PanicIfError(err)

	task.Title = request.Title
	task.Status = request.Status
	task.Priority = request.Priority
	task.Effort = request.Effort
	task.DifficultyLevel = request.DifficultyLevel
	task.Deliverable = request.Deliverable
	task.Bottleneck = request.Bottleneck
	task.UpdatedAt = time.Now()

	result, err := service.TaskRepository.Update(ctx, task)
	helper.PanicIfError(err)

	return helper.ToTaskResponse(result)
}

func (service *TaskServiceImpl) Delete(ctx context.Context, taskId uuid.UUID) {
	err := service.TaskRepository.Delete(ctx, taskId)
	helper.PanicIfError(err)
}

func (service *TaskServiceImpl) FindById(ctx context.Context, taskId uuid.UUID) web.TaskResponse {
	task, err := service.TaskRepository.FindById(ctx, taskId)
	helper.PanicIfError(err)

	return helper.ToTaskResponse(task)
}

func (service *TaskServiceImpl) FindByProjectId(ctx context.Context, projectId uuid.UUID) []web.TaskResponse {
	tasks, err := service.TaskRepository.FindByProjectId(ctx, projectId)
	helper.PanicIfError(err)

	return helper.ToTaskResponses(tasks)
}

func (service *TaskServiceImpl) FindAll(ctx context.Context) []web.TaskResponse {
	tasks, err := service.TaskRepository.FindAll(ctx)
	helper.PanicIfError(err)

	return helper.ToTaskResponses(tasks)
}
