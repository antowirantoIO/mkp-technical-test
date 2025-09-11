package request

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	Username  string  `json:"username" validate:"required,max=100"`
	Email     string  `json:"email" validate:"required,email,max=255"`
	Password  string  `json:"password" validate:"required,min=8,max=100"`
	FirstName *string `json:"first_name" validate:"omitempty,max=100"`
	LastName  *string `json:"last_name" validate:"omitempty,max=100"`
	Phone     *string `json:"phone" validate:"omitempty,max=20"`
}

type UpdateUserRequest struct {
	ID        string  `json:"-" validate:"required,max=100,uuid"`
	Username  *string `json:"username" validate:"omitempty,max=100"`
	Email     *string `json:"email" validate:"omitempty,email,max=255"`
	Password  *string `json:"password" validate:"omitempty,min=8,max=100"`
	FirstName *string `json:"first_name" validate:"omitempty,max=100"`
	LastName  *string `json:"last_name" validate:"omitempty,max=100"`
	Phone     *string `json:"phone" validate:"omitempty,max=20"`
	Avatar    *string `json:"avatar" validate:"omitempty,url,max=500"`
	IsActive  *bool   `json:"is_active"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type GetUserRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteUserRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type ListUserRequest struct {
	Page       int     `json:"page" validate:"min=1"`
	Size       int     `json:"size" validate:"min=1,max=100"`
	IsActive   *bool   `json:"is_active"`
	IsVerified *bool   `json:"is_verified"`
	Username   *string `json:"username"`
	Email      *string `json:"email"`
}

type ChangePasswordRequest struct {
	ID              string `json:"-" validate:"required,max=100,uuid"`
	CurrentPassword string `json:"current_password" validate:"required,max=100"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=100"`
}

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required,max=255"`
}

type ResendVerificationRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email,max=255"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required,max=255"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=100"`
}

type AssignRolesRequest struct {
	UserID  string   `json:"-" validate:"required,max=100,uuid"`
	RoleIDs []string `json:"role_ids" validate:"required,dive,uuid"`
}

type RemoveRolesRequest struct {
	UserID  string   `json:"-" validate:"required,max=100,uuid"`
	RoleIDs []string `json:"role_ids" validate:"required,dive,uuid"`
}

type GetUsersByRoleIdRequest struct {
	RoleID string `json:"-" validate:"required,max=100,uuid"`
}
