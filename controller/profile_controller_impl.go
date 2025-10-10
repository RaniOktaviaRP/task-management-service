package controller

import (
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"task-management/helper"
	"task-management/model/web"
	"task-management/service"
)

type ProfileControllerImpl struct {
	ProfileService service.ProfileService
}

func NewProfileController(profileService service.ProfileService) ProfileController {
	return &ProfileControllerImpl{
		ProfileService: profileService,
	}
}

func (c *ProfileControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request web.ProfileCreateRequest
	if err := helper.ReadFromRequestBody(r, &request); err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	response := c.ProfileService.Create(r.Context(), request)
	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProfileControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request web.ProfileUpdateRequest
	if err := helper.ReadFromRequestBody(r, &request); err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	profileId, err := uuid.Parse(ps.ByName("profileId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid profile ID",
		})
		return
	}
	request.Id = profileId

	response := c.ProfileService.Update(r.Context(), request)
	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProfileControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileId, err := uuid.Parse(ps.ByName("profileId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid profile ID",
		})
		return
	}

	c.ProfileService.Delete(r.Context(), profileId)
	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
	})
}

func (c *ProfileControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	profileId, err := uuid.Parse(ps.ByName("profileId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid profile ID",
		})
		return
	}

	response := c.ProfileService.FindById(r.Context(), profileId)
	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProfileControllerImpl) FindByUserId(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId, err := uuid.Parse(ps.ByName("userId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		})
		return
	}

	response := c.ProfileService.FindByUserId(r.Context(), userId)
	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (c *ProfileControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	responses := c.ProfileService.FindAll(r.Context())
	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}