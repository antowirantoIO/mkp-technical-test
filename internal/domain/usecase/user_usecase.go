package usecase

import (
	"context"
	"mkp-boarding-test/internal/delivery/http/dto/request"
	"mkp-boarding-test/internal/delivery/http/dto/response"
)

type UserUseCase interface {
	Verify(ctx context.Context, request *request.VerifyUserRequest) (*response.UserResponse, error)
	Create(ctx context.Context, request *request.RegisterUserRequest) (*response.UserResponse, error)
	Login(ctx context.Context, request *request.LoginUserRequest) (*response.UserResponse, error)
	Current(ctx context.Context, request *request.GetUserRequest) (*response.UserResponse, error)
	Logout(ctx context.Context, request *request.LogoutUserRequest) (bool, error)
	Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error)
}