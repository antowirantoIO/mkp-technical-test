package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
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
