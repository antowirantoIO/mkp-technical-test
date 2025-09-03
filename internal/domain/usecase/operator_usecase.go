package usecase

import (
	"context"
	"mkp-boarding-test/internal/model"
)

type OperatorUseCase interface {
	Create(ctx context.Context, request *model.CreateOperatorRequest) (*model.OperatorResponse, error)
	Update(ctx context.Context, request *model.UpdateOperatorRequest) (*model.OperatorResponse, error)
	Get(ctx context.Context, request *model.GetOperatorRequest) (*model.OperatorResponse, error)
	Delete(ctx context.Context, request *model.DeleteOperatorRequest) error
	List(ctx context.Context, request *model.ListOperatorRequest) (*model.WebResponse[[]model.OperatorResponse], error)
}