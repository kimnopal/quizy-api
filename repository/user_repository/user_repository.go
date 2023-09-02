package repository

import (
	"context"
	"database/sql"
	"quizy-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, userId int64)
	FindById(ctx context.Context, tx *sql.Tx, userId int64) (domain.User, error)
}
