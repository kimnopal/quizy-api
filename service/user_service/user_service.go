package user_service

import (
	"context"
	"quizy-api/model/web/user_dto"
)

type UserService interface {
	Create(ctx context.Context, request user_dto.UserCreateRequest) user_dto.UserResponse
	Update(ctx context.Context, request user_dto.UserUpdateRequest) user_dto.UserResponse
	Delete(ctx context.Context, userId int64)
}
