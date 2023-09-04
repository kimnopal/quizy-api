package domain

import "time"

type Quiz struct {
	Id             int64     `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Code           string    `json:"code"`
	InvitationCode string    `json:"invitation_code"`
	IsActive       bool      `json:"is_active"`
	UserId         int64     `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
