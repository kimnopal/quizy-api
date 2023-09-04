package helper

import (
	"quizy-api/model/domain"
	"quizy-api/model/web/role_dto"
	"quizy-api/model/web/user_dto"
)

func ToRoleResponse(role domain.Role) role_dto.RoleResponse {
	return role_dto.RoleResponse{
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
