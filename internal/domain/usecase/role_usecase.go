package usecase

import (
	"context"
	"mkp-boarding-test/internal/model"
)

type RoleUseCase interface {
	Create(ctx context.Context, request *model.CreateRoleRequest) (*model.RoleResponse, error)
	Update(ctx context.Context, request *model.UpdateRoleRequest) (*model.RoleResponse, error)
	Get(ctx context.Context, request *model.GetRoleRequest) (*model.RoleResponse, error)
	Delete(ctx context.Context, request *model.DeleteRoleRequest) error
	List(ctx context.Context, request *model.ListRoleRequest) (*model.WebResponse[[]model.RoleResponse], error)
	AssignPermissions(ctx context.Context, request *model.AssignPermissionsRequest) error
	RemovePermissions(ctx context.Context, request *model.RemovePermissionsRequest) error
}