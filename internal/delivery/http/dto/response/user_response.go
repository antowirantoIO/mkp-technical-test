package response

type UserResponse struct {
	ID                string  `json:"id"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	FirstName         *string `json:"first_name"`
	LastName          *string `json:"last_name"`
	Phone             *string `json:"phone"`
	Avatar            *string `json:"avatar"`
	IsActive          bool    `json:"is_active"`
	IsVerified        bool    `json:"is_verified"`
	LastLoginAt       *int64  `json:"last_login_at"`
	EmailVerifiedAt   *int64  `json:"email_verified_at"`
	PasswordChangedAt *int64  `json:"password_changed_at"`
	Token             string  `json:"token,omitempty"`
	CreatedAt         int64   `json:"created_at"`
	UpdatedAt         int64   `json:"updated_at"`
}