package model

type PermissionResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Description *string `json:"description"`
	Resource    string  `json:"resource"`
	Action      string  `json:"action"`
	IsActive    bool    `json:"is_active"`
	IsSystem    bool    `json:"is_system"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type CreatePermissionRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	DisplayName string  `json:"display_name" validate:"required,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	Resource    string  `json:"resource" validate:"required,max=100"`
	Action      string  `json:"action" validate:"required,max=100"`
	IsActive    *bool   `json:"is_active"`
}

type UpdatePermissionRequest struct {
	ID          string  `json:"-" validate:"required,max=100,uuid"`
	Name        *string `json:"name" validate:"omitempty,max=100"`
	DisplayName *string `json:"display_name" validate:"omitempty,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	Resource    *string `json:"resource" validate:"omitempty,max=100"`
	Action      *string `json:"action" validate:"omitempty,max=100"`
	IsActive    *bool   `json:"is_active"`
}

type GetPermissionRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeletePermissionRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type ListPermissionRequest struct {
	Page     int     `json:"page" validate:"min=1"`
	Size     int     `json:"size" validate:"min=1,max=100"`
	IsActive *bool   `json:"is_active"`
	Name     *string `json:"name"`
	Resource *string `json:"resource"`
	Action   *string `json:"action"`
}