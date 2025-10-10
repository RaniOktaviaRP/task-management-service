package helper

import (
	"task-management/model/domain"
	"task-management/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:           user.Id,          
		Email:        user.Email,
		Role:         user.Role,
	}
}
