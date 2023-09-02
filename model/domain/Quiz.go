package domain

import "time"

type Quiz struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Code           string    `json:"code"`
	InvitationCode string    `json:"invitation_code"`
	IsActive       bool      `json:"is_active"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
