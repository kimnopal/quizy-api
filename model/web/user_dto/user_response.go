package user_dto

type UserResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleId int64  `json:"role_id"`
}
