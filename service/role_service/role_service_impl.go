package service

import (
	"context"
	"database/sql"
	"quizy-api/exception"
	"quizy-api/helper"
	"quizy-api/model/domain"
	web_role "quizy-api/model/web/role"
	repository "quizy-api/repository/role_repository"

	"github.com/go-playground/validator/v10"
)

type RoleServiceImpl struct {
	DB         *sql.DB
	repository repository.RoleRepository
	validate   *validator.Validate
}

func NewRoleService(DB *sql.DB, repository repository.RoleRepository, validate *validator.Validate) RoleService {
	return &RoleServiceImpl{
		DB:         DB,
		repository: repository,
		validate:   validate,
	}
}

func (service *RoleServiceImpl) Create(ctx context.Context, request web_role.RoleCreateRequest) web_role.RoleResponse {
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

func (service *RoleServiceImpl) Update(ctx context.Context, request web_role.RoleUpdateRequest) web_role.RoleResponse {
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
