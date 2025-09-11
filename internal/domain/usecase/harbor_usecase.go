package usecase

import (
	"context"
	"mkp-boarding-test/internal/model"
)

type HarborUseCase interface {
	Create(ctx context.Context, request *model.CreateHarborRequest) (*model.HarborResponse, error)
	Update(ctx context.Context, request *model.UpdateHarborRequest) (*model.HarborResponse, error)
	Get(ctx context.Context, request *model.GetHarborRequest) (*model.HarborResponse, error)
	Delete(ctx context.Context, request *model.DeleteHarborRequest) error
	List(ctx context.Context, request *model.ListHarborRequest, userId string) (*model.WebResponse[[]model.HarborResponse], error)
}
