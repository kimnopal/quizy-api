package user_service

import (
	"context"
	"database/sql"
	"quizy-api/helper"
	"quizy-api/model/domain"
	"quizy-api/model/web/user_dto"
	"quizy-api/repository/user_repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	DB         *sql.DB
	repository user_repository.UserRepository
	validate   *validator.Validate
}

func NewUserService(DB *sql.DB, repository user_repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		DB:         DB,
		repository: repository,
		validate:   validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request user_dto.UserCreateRequest) user_dto.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		RoleId:   request.RoleId,
	}
	user = service.repository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user_dto.UserUpdateRequest) user_dto.UserResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.repository.FindById(ctx, tx, request.Id)
	helper.PanicIfNotFound(err)

	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password
	user.RoleId = request.RoleId

	user = service.repository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int64) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.repository.FindById(ctx, tx, userId)
	helper.PanicIfNotFound(err)

	service.repository.Delete(ctx, tx, user.Id)

}
