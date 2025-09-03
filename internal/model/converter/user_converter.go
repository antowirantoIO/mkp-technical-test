package converter

import (
	"mkp-boarding-test/internal/delivery/http/dto/response"
	"mkp-boarding-test/internal/domain/entity"
)

func UserToResponse(user *entity.User) *response.UserResponse {
	return &response.UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		FirstName:         &user.FirstName,
		LastName:          &user.LastName,
		Phone:             user.Phone,
		Avatar:            user.Avatar,
		IsActive:          user.IsActive,
		IsVerified:        user.IsVerified,
		LastLoginAt:       user.LastLoginAt,
		EmailVerifiedAt:   user.EmailVerifiedAt,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}
