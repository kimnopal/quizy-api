package helper

import (
	"quizy-api/model/domain"
	web_role "quizy-api/model/web/role"
)

func ToRoleResponse(role domain.Role) web_role.RoleResponse {
	return web_role.RoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}
}
