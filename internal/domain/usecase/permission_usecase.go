package usecase

import (
	"context"
	"mkp-boarding-test/internal/model"
)

type PermissionUseCase interface {
	Create(ctx context.Context, request *model.CreatePermissionRequest) (*model.PermissionResponse, error)
	Update(ctx context.Context, request *model.UpdatePermissionRequest) (*model.PermissionResponse, error)
	Get(ctx context.Context, request *model.GetPermissionRequest) (*model.PermissionResponse, error)
	Delete(ctx context.Context, request *model.DeletePermissionRequest) error
	List(ctx context.Context, request *model.ListPermissionRequest) (*model.WebResponse[[]model.PermissionResponse], error)
}