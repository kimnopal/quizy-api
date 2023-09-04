package role_dto

type RoleUpdateRequest struct {
	Id   int64  `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=255" json:"name"`
}
