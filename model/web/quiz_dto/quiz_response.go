package quiz_dto

type QuizResponse struct {
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Code           string `json:"code"`
	InvitationCode string `json:"invitation_code"`
	IsActive       bool   `json:"is_active"`
}
