package user_dto

type UserUpdateRequest struct {
	Id       int64  `json:"id" validate:"number"`
	Name     string `json:"name" validate:"min=1,max=255"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8"`
	RoleId   int64  `json:"role_id" validate:"number"`
}
