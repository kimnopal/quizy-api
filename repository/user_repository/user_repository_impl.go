package user_repository

import (
	"context"
	"database/sql"
	"errors"
	"quizy-api/helper"
	"quizy-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "INSERT INTO user (name, email, password, role_id) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, user.RoleId)
	helper.PanicIfError(err)

	index, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = index

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "UPDATE user SET name = ?, email = ?, password = ?, role_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, user.Name, user.Email, user.Password, user.RoleId, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int64) {
	sql := "DELETE FROM user WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, userId)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int64) (domain.User, error) {
	sql := "SELECT id, name, email, password, role_id FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.RoleId)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
