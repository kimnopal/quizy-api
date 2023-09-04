package quiz_repository

import (
	"context"
	"database/sql"
	"errors"
	"quizy-api/helper"
	"quizy-api/model/domain"
)

type QuizRepositoryImpl struct {
}

func (repository *QuizRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, quiz domain.Quiz) domain.Quiz {
	sql := "INSERT INTO quiz (title, description, code, invitation_code, is_active, user_id) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, quiz.Title, quiz.Description, quiz.Code, quiz.InvitationCode, quiz.IsActive, quiz.UserId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	quiz.Id = id
	return quiz
}

func (repository *QuizRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, quiz domain.Quiz) domain.Quiz {
	sql := "UPDATE quiz SET (title, description, invitation_code, is_active) WHERE code = ?"
	_, err := tx.ExecContext(ctx, sql, quiz.Code)
	helper.PanicIfError(err)

	return quiz
}

func (repository *QuizRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, quizCode string) {
	sql := "DELETE FROM quiz WHERE code = ?"
	_, err := tx.ExecContext(ctx, sql, quizCode)
	helper.PanicIfError(err)
}

func (repository *QuizRepositoryImpl) FindByCode(ctx context.Context, tx *sql.Tx, quizCode string) (domain.Quiz, error) {
	sql := "SELECT id, title, description, code, invitation_code, is_active FROM quiz WHERE code = ?"
	rows, err := tx.QueryContext(ctx, sql, quizCode)
	helper.PanicIfError(err)
	defer rows.Close()

	quiz := domain.Quiz{}
	if rows.Next() {
		err := rows.Scan(&quiz.Id, &quiz.Title, &quiz.Description, &quiz.Code, &quiz.InvitationCode, &quiz.IsActive)
		helper.PanicIfError(err)
		return quiz, nil
	} else {
		return quiz, errors.New("quiz not found")
	}
}

func (repository *QuizRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Quiz {
	sql := "SELECT id, title, description, code, invitation_code, is_active FROM quiz"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	quizzes := []domain.Quiz{}
	for rows.Next() {
		quiz := domain.Quiz{}
		err = rows.Scan(&quiz.Id, &quiz.Title, &quiz.Description, &quiz.Code, &quiz.InvitationCode, &quiz.IsActive)
		helper.PanicIfError(err)

		quizzes = append(quizzes, quiz)
	}

	return quizzes
}
