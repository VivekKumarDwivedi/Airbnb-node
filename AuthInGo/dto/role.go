package dto

type CreateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=255"`
}
type UpdateRoleRequestDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=255"`
}

type AddPermissionToRoleRequestDTO struct {
	RoleID       int64 `json:"role_id" validate:"required"`
	PermissionID int64 `json:"permission_id" validate:"required"`
}
type RemovePermissionFromRoleRequestDTO struct {
	RoleID       int64 `json:"role_id" validate:"required"`
	PermissionID int64 `json:"permission_id" validate:"required"`
}
