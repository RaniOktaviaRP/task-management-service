package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"task-management/helper"
	"task-management/model/web"
	"task-management/service"
)

// UserControllerImpl implements UserController interface
type UserControllerImpl struct {
	UserService service.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user with email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body web.UserRegisterRequest true "User payload"
// @Success 200 {object} web.WebResponse{data=web.UserResponse}
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /users [post]
func (c *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req web.UserRegisterRequest
	if err := helper.ReadFromRequestBody(r, &req); err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	res, err := c.UserService.Register(r.Context(), req)
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Users
// @Accept json
// @Produce json
// @Param request body web.UserLoginRequest true "User login payload"
// @Success 200 {object} web.WebResponse{data=web.TokenResponse}
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Router /login [post]
func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req web.UserLoginRequest
	if err := helper.ReadFromRequestBody(r, &req); err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	token, refreshToken, err := c.UserService.Login(r.Context(), req)
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   401,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: web.TokenResponse{
			Token:        token,
			RefreshToken: refreshToken,
		},
	})
}

// Refresh token godoc
// @Summary Refresh JWT token
// @Description Generate a new JWT token using refresh token
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param X-Refresh-Token header string true "Refresh token"
// @Success 200 {object} web.WebResponse{data=web.TokenResponse}
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 401 {object} web.WebResponse "Unauthorized"
// @Router /refresh [post]
func (c *UserControllerImpl) Refresh(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	refreshToken := r.Header.Get("X-Refresh-Token")
	if refreshToken == "" {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "refresh token tidak ditemukan di header",
		})
		return
	}

	newToken, newRefreshToken, err := c.UserService.Refresh(r.Context(), refreshToken)
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   401,
			Status: "UNAUTHORIZED",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data: web.TokenResponse{
			Token:        newToken,
			RefreshToken: newRefreshToken,
		},
	})
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate refresh token
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param X-Refresh-Token header string true "Refresh token"
// @Success 200 {object} web.WebResponse "Logout berhasil"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /logout [post]
func (c *UserControllerImpl) Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	refreshToken := r.Header.Get("X-Refresh-Token")
	if refreshToken == "" {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "refresh token tidak ditemukan di header",
		})
		return
	}

	err := c.UserService.Logout(r.Context(), refreshToken)
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "logout berhasil",
	})
}

// Update godoc
// @Summary Update a user
// @Description Update an existing user by UUID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID (UUID)"
// @Param user body web.UserUpdateRequest true "User payload"
// @Success 200 {object} web.WebResponse{data=web.UserResponse}
// @Failure 400 {object} web.WebResponse "Invalid UUID or Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Security BearerAuth
// @Router /users/{userId} [put]
func (c *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req web.UserUpdateRequest
	if err := helper.ReadFromRequestBody(r, &req); err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
		return
	}

	uid, err := uuid.Parse(ps.ByName("userId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "INVALID UUID",
			Data:   err.Error(),
		})
		return
	}
	req.Id = uid

	res, err := c.UserService.Update(r.Context(), req)
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}

// Delete godoc
// @Summary Delete a user
// @Description Soft delete a user by UUID
// @Tags Users
// @Produce json
// @Param userId path string true "User ID (UUID)"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse "Invalid UUID"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Security BearerAuth
// @Router /users/{userId} [delete]
func (c *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, err := uuid.Parse(ps.ByName("userId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "INVALID UUID",
			Data:   err.Error(),
		})
		return
	}

	if err := c.UserService.Delete(r.Context(), uid); err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
	})
}

// FindById godoc
// @Summary Get user by ID
// @Description Get a single user by UUID
// @Tags Users
// @Produce json
// @Param userId path string true "User ID (UUID)"
// @Success 200 {object} web.WebResponse{data=web.UserResponse}
// @Failure 400 {object} web.WebResponse "Invalid UUID"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Security BearerAuth
// @Router /users/{userId} [get]
func (c *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, err := uuid.Parse(ps.ByName("userId"))
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   400,
			Status: "INVALID UUID",
			Data:   err.Error(),
		})
		return
	}

	res, err := c.UserService.FindById(r.Context(), uid)
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}

// FindAll godoc
// @Summary Get all users
// @Description Get all users without pagination
// @Tags Users
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.UserResponse}
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Security BearerAuth
// @Router /users [get]
func (c *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := c.UserService.FindAll(r.Context())
	if err != nil {
		helper.WriteToResponseBody(w, web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   res,
	})
}
