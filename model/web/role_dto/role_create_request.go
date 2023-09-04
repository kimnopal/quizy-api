package role_dto

type RoleCreateRequest struct {
	Name string `validate:"required,min=1,max=255" json:"name"`
}
