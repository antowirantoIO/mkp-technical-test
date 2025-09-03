package usecase

import (
	"context"
	"mkp-boarding-test/internal/model"
)

type ShipUseCase interface {
	Create(ctx context.Context, request *model.CreateShipRequest) (*model.ShipResponse, error)
	Update(ctx context.Context, request *model.UpdateShipRequest) (*model.ShipResponse, error)
	Get(ctx context.Context, request *model.GetShipRequest) (*model.ShipResponse, error)
	Delete(ctx context.Context, request *model.DeleteShipRequest) error
	List(ctx context.Context, request *model.ListShipRequest) (*model.WebResponse[[]model.ShipResponse], error)
}