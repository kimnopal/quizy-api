package role_repository

import (
	"context"
	"database/sql"
	"errors"
	"quizy-api/helper"
	"quizy-api/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	sql := "INSERT INTO role (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, sql, role.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = id
	return role
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	sql := "UPDATE role SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, role.Name, role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, roleId int64) {
	sql := "DELETE FROM role WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, roleId)
	helper.PanicIfError(err)
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId int64) (domain.Role, error) {
	sql := "SELECT id, name FROM role WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, roleId)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}
	if rows.Next() {
		err = rows.Scan(&role.Id, &role.Name)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("role not found")
	}
}
