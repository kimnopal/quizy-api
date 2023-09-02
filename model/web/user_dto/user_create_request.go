package user_dto

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	RoleId   int64  `json:"role_id" validate:"required,number"`
}
