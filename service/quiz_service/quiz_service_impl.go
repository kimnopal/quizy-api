package quiz_service

import (
	"context"
	"database/sql"
	"quizy-api/helper"
	"quizy-api/model/domain"
	"quizy-api/model/web/quiz_dto"
	"quizy-api/repository/quiz_repository"

	"github.com/go-playground/validator/v10"
)

type QuizServiceImpl struct {
	DB         *sql.DB
	repository quiz_repository.QuizRepository
	validate   *validator.Validate
}

func (service *QuizServiceImpl) Create(ctx context.Context, request quiz_dto.QuizCreateRequest) quiz_dto.QuizResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	quiz := domain.Quiz{
		Title:          request.Title,
		Description:    request.Description,
		InvitationCode: request.InvitationCode,
		IsActive:       request.IsActive,
	}

	quiz = service.repository.Save(ctx, tx, quiz)

	return helper.ToQuizResponse(quiz)
}

func (service *QuizServiceImpl) Update(ctx context.Context, request quiz_dto.QuizUpdateRequest) quiz_dto.QuizResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	quiz, err := service.repository.FindByCode(ctx, tx, request.Code)
	helper.PanicIfNotFound(err)

	quiz.Title = request.Title
	quiz.Description = request.Description
	quiz.InvitationCode = request.InvitationCode
	quiz.IsActive = request.IsActive

	quiz = service.repository.Update(ctx, tx, quiz)

	return helper.ToQuizResponse(quiz)
}

func (service *QuizServiceImpl) Delete(ctx context.Context, quizCode string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.repository.FindByCode(ctx, tx, quizCode)
	helper.PanicIfNotFound(err)

	service.repository.Delete(ctx, tx, quizCode)
}

func (service *QuizServiceImpl) FindByCode(ctx context.Context, quizCode string) quiz_dto.QuizResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	quiz, err := service.repository.FindByCode(ctx, tx, quizCode)
	helper.PanicIfNotFound(err)

	return helper.ToQuizResponse(quiz)
}

func (service *QuizServiceImpl) FindAll(ctx context.Context) []quiz_dto.QuizResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	quizzes := service.repository.FindAll(ctx, tx)

	return helper.ToQuizResponses(quizzes)
}
