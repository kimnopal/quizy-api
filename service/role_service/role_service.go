package role_service

import (
	"context"
	"quizy-api/model/web/role_dto"
)

type RoleService interface {
	Create(ctx context.Context, request role_dto.RoleCreateRequest) role_dto.RoleResponse
	Update(ctx context.Context, request role_dto.RoleUpdateRequest) role_dto.RoleResponse
	Delete(ctx context.Context, roleId int64)
}
