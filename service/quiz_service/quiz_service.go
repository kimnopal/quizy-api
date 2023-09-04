package quiz_service

import (
	"context"
	"quizy-api/model/web/quiz_dto"
)

type QuizService interface {
	Create(ctx context.Context, request quiz_dto.QuizCreateRequest) quiz_dto.QuizResponse
	Update(ctx context.Context, request quiz_dto.QuizUpdateRequest) quiz_dto.QuizResponse
	Delete(ctx context.Context, quizCode string)
	FindByCode(ctx context.Context, quizCode string) quiz_dto.QuizResponse
	FindAll(ctx context.Context) []quiz_dto.QuizResponse
}
