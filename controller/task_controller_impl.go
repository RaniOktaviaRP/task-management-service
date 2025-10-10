package controller

import (
	"net/http"
	"task-management/helper"
	"task-management/model/web"
	"task-management/service"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type TaskControllerImpl struct {
	TaskService service.TaskService
}

func NewTaskController(taskService service.TaskService) TaskController {
	return &TaskControllerImpl{
		TaskService: taskService,
	}
}

// Create godoc
// @Summary Create a new task
// @Description Create a new task with the provided data
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body web.TaskCreateRequest true "Task data"
// @Success 200 {object} web.TaskResponse
// @Security BearerAuth
// @Router /tasks [post]
func (controller *TaskControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskCreateRequest := web.TaskCreateRequest{}
	helper.ReadFromRequestBody(request, &taskCreateRequest)

	taskResponse := controller.TaskService.Create(request.Context(), taskCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update godoc
// @Summary Update a task
// @Description Update a task with the provided data
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body web.TaskUpdateRequest true "Task data"
// @Success 200 {object} web.TaskResponse
// @Security BearerAuth
// @Router /tasks/{id} [put]
func (controller *TaskControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskUpdateRequest := web.TaskUpdateRequest{}
	helper.ReadFromRequestBody(request, &taskUpdateRequest)

	taskId, err := uuid.Parse(params.ByName("id"))
	helper.PanicIfError(err)

	taskResponse := controller.TaskService.Update(request.Context(), taskId, taskUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete godoc
// @Summary Delete a task
// @Description Delete a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} web.WebResponse
// @Security BearerAuth
// @Router /tasks/{id} [delete]
func (controller *TaskControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId, err := uuid.Parse(params.ByName("id"))
	helper.PanicIfError(err)

	controller.TaskService.Delete(request.Context(), taskId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById godoc
// @Summary Get a task by ID
// @Description Get a task by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} web.TaskResponse
// @Security BearerAuth
// @Router /tasks/{id} [get]
func (controller *TaskControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskId, err := uuid.Parse(params.ByName("id"))
	helper.PanicIfError(err)

	taskResponse := controller.TaskService.FindById(request.Context(), taskId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindByProjectId godoc
// @Summary Get tasks by project ID
// @Description Get all tasks for a specific project
// @Tags tasks
// @Accept json
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200 {array} web.TaskResponse
// @Security BearerAuth
// @Router /tasks/project/{projectId} [get]
func (controller *TaskControllerImpl) FindByProjectId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	projectId, err := uuid.Parse(params.ByName("projectId"))
	helper.PanicIfError(err)

	taskResponses := controller.TaskService.FindByProjectId(request.Context(), projectId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll godoc
// @Summary Get all tasks
// @Description Get all tasks in the system
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {array} web.TaskResponse
// @Security BearerAuth
// @Router /tasks [get]
func (controller *TaskControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	taskResponses := controller.TaskService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   taskResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}