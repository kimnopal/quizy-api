package service

import (
	"context"
	web_role "quizy-api/model/web/role"
)

type RoleService interface {
	Create(ctx context.Context, request web_role.RoleCreateRequest) web_role.RoleResponse
	Update(ctx context.Context, request web_role.RoleUpdateRequest) web_role.RoleResponse
	Delete(ctx context.Context, roleId int64)
}
