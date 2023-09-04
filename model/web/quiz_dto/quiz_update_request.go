package quiz_dto

type QuizUpdateRequest struct {
	Code           string `json:"code"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	InvitationCode string `json:"invitation"`
	IsActive       bool   `json:"is_active"`
}
