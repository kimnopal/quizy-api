package helper

import (
	"quizy-api/model/domain"
	web_role "quizy-api/model/web/role"
	"quizy-api/model/web/user_dto"
)

func ToRoleResponse(role domain.Role) web_role.RoleResponse {
	return web_role.RoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}
}

func ToUserResponse(user domain.User) user_dto.UserResponse {
	return user_dto.UserResponse{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		RoleId: user.RoleId,
	}
}
