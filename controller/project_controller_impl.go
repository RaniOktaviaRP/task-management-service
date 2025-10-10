package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"task-management/helper"
	"task-management/model/web"
	"task-management/service"
)

// ProjectControllerImpl adalah implementasi dari ProjectController
type ProjectControllerImpl struct {
	ProjectService service.ProjectService
}

// NewProjectController membuat instance ProjectController baru
func NewProjectController(projectService service.ProjectService) ProjectController {
	return &ProjectControllerImpl{
		ProjectService: projectService,
	}
}

// @Summary Create new project
// @Description Create new project with the input payload
// @Tags projects
// @Accept json
// @Produce json
// @Param project body web.ProjectCreateRequest true "Create project request"
// @Success 200 {object} web.ProjectResponse
// @Security BearerAuth
// @Router /projects [post]
func (controller *ProjectControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	projectCreateRequest := web.ProjectCreateRequest{}
	helper.ReadFromRequestBody(request, &projectCreateRequest)

	projectResponse := controller.ProjectService.Create(request.Context(), projectCreateRequest)
	webResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   projectResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Update project
// @Description Update project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body web.ProjectUpdateRequest true "Update project request"
// @Success 200 {object} web.ProjectResponse
// @Security BearerAuth
// @Router /projects/{id} [put]
func (controller *ProjectControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	projectUpdateRequest := web.ProjectUpdateRequest{}
	helper.ReadFromRequestBody(request, &projectUpdateRequest)

	projectId, err := uuid.Parse(params.ByName("id"))
	helper.PanicIfError(err)
	projectUpdateRequest.Id = projectId

	projectResponse := controller.ProjectService.Update(request.Context(), projectUpdateRequest)
	webResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   projectResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Delete project
// @Description Delete project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} map[string]interface{} "response with code and status"
// @Security BearerAuth
// @Router /projects/{id} [delete]
func (controller *ProjectControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	projectId, err := uuid.Parse(params.ByName("id"))
	helper.PanicIfError(err)

	controller.ProjectService.Delete(request.Context(), projectId)
	webResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Get project by ID
// @Description Get project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} web.ProjectResponse
// @Security BearerAuth
// @Router /projects/{id} [get]
func (controller *ProjectControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	projectId, err := uuid.Parse(params.ByName("id"))
	helper.PanicIfError(err)

	projectResponse := controller.ProjectService.FindById(request.Context(), projectId)
	webResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   projectResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Get projects by user ID
// @Description Get all projects for a specific user
// @Tags projects
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {array} web.ProjectResponse
// @Security BearerAuth
// @Router /projects/user/{userId} [get]
func (controller *ProjectControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId, err := uuid.Parse(params.ByName("userId"))
	helper.PanicIfError(err)

	projectResponses := controller.ProjectService.FindByUserId(request.Context(), userId)
	webResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   projectResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// @Summary Get all projects
// @Description Get all projects
// @Tags projects
// @Accept json
// @Produce json
// @Success 200 {array} web.ProjectResponse
// @Security BearerAuth
// @Router /projects [get]
func (controller *ProjectControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	projectResponses := controller.ProjectService.FindAll(request.Context())
	webResponse := helper.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   projectResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
