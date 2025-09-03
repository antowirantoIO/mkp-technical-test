package usecase

import (
	"context"
	"mkp-boarding-test/internal/model"
)

type AuthUseCase interface {
	Register(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error)
	Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error)
	Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error)
	VerifyToken(ctx context.Context, request *model.VerifyUserRequest) (*model.UserResponse, error)
}