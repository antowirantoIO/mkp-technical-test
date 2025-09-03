package converter

import (
	"mkp-boarding-test/internal/entity"
	"mkp-boarding-test/internal/model"
)

func PermissionToResponse(permission *entity.Permission) *model.PermissionResponse {
	return &model.PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		DisplayName: permission.DisplayName,
		Description: permission.Description,
		Resource:    permission.Resource,
		Action:      permission.Action,
		IsActive:    permission.IsActive,
		IsSystem:    permission.IsSystem,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}
