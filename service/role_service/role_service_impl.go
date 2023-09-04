package role_service

import (
	"context"
	"database/sql"
	"quizy-api/exception"
	"quizy-api/helper"
	"quizy-api/model/domain"
	"quizy-api/model/web/role_dto"
	"quizy-api/repository/role_repository"

	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	DB         *sql.DB
	repository role_repository.RoleRepository
	validate   *validator.Validate
}

func NewRoleService(DB *sql.DB, repository role_repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		DB:         DB,
		repository: repository,
		validate:   validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request role_dto.RoleCreateRequest) role_dto.RoleResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := domain.Role{
		Name: request.Name,
	}

	role = service.repository.Save(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Update(ctx context.Context, request role_dto.RoleUpdateRequest) role_dto.RoleResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.repository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	role.Name = request.Name

	role = service.repository.Update(ctx, tx, role)

	return helper.ToRoleResponse(role)
}

func (service *RoleServiceImpl) Delete(ctx context.Context, roleId int64) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.repository.FindById(ctx, tx, roleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.repository.Delete(ctx, tx, role.Id)
}
