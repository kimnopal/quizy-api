package role_repository

import (
	"context"
	"database/sql"
	"quizy-api/model/domain"
)

type RoleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Delete(ctx context.Context, tx *sql.Tx, roleId int64)
	FindById(ctx context.Context, tx *sql.Tx, roleId int64) (domain.Role, error)
}
