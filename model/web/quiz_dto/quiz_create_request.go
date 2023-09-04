package quiz_dto

type QuizCreateRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	InvitationCode string `json:"invitation_code"`
	IsActive       bool   `json:"is_active"`
}
