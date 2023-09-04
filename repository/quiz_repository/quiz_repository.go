package quiz_repository

import (
	"context"
	"database/sql"
	"quizy-api/model/domain"
)

type QuizRepository interface {
	Save(ctx context.Context, tx *sql.Tx, quiz domain.Quiz) domain.Quiz
	Update(ctx context.Context, tx *sql.Tx, quiz domain.Quiz) domain.Quiz
	Delete(ctx context.Context, tx *sql.Tx, quizCode string)
	FindByCode(ctx context.Context, tx *sql.Tx, quizCode string) (domain.Quiz, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Quiz
}
