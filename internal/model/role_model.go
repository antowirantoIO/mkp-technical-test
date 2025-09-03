package model

type RoleResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Description *string `json:"description"`
	IsActive    bool    `json:"is_active"`
	IsSystem    bool    `json:"is_system"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type CreateRoleRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	DisplayName string  `json:"display_name" validate:"required,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	IsActive    *bool   `json:"is_active"`
}

type UpdateRoleRequest struct {
	ID          string  `json:"-" validate:"required,max=100,uuid"`
	Name        *string `json:"name" validate:"omitempty,max=100"`
	DisplayName *string `json:"display_name" validate:"omitempty,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	IsActive    *bool   `json:"is_active"`
}

type GetRoleRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteRoleRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type ListRoleRequest struct {
	Page     int     `json:"page" validate:"min=1"`
	Size     int     `json:"size" validate:"min=1,max=100"`
	IsActive *bool   `json:"is_active"`
	Name     *string `json:"name"`
}

type AssignPermissionsRequest struct {
	RoleID        string   `json:"-" validate:"required,max=100,uuid"`
	PermissionIDs []string `json:"permission_ids" validate:"required,dive,uuid"`
}

type RemovePermissionsRequest struct {
	RoleID        string   `json:"-" validate:"required,max=100,uuid"`
	PermissionIDs []string `json:"permission_ids" validate:"required,dive,uuid"`
}
