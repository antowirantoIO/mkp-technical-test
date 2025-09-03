package converter

import (
	"mkp-boarding-test/internal/domain/entity"
	"mkp-boarding-test/internal/model"
)

func RoleToResponse(role *entity.Role) *model.RoleResponse {
	return &model.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		DisplayName: role.DisplayName,
		Description: role.Description,
		IsActive:    role.IsActive,
		IsSystem:    role.IsSystem,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}
