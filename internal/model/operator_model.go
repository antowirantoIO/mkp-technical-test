package model

type OperatorResponse struct {
	ID             string  `json:"id"`
	UserID         string  `json:"user_id"`
	OperatorCode   string  `json:"operator_code"`
	CompanyName    string  `json:"company_name"`
	LicenseNumber  string  `json:"license_number"`
	ContactPerson  string  `json:"contact_person"`
	ContactPhone   string  `json:"contact_phone"`
	ContactEmail   string  `json:"contact_email"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	Province       string  `json:"province"`
	Country        string  `json:"country"`
	PostalCode     string  `json:"postal_code"`
	Website        *string `json:"website"`
	OperatorType   string  `json:"operator_type"`
	Status         string  `json:"status"`
	EstablishedAt  *int64  `json:"established_at"`
	LicenseExpiry  *int64  `json:"license_expiry"`
	IsActive       bool    `json:"is_active"`
	Notes          *string `json:"notes"`
	CreatedAt      int64   `json:"created_at"`
	UpdatedAt      int64   `json:"updated_at"`
}

type CreateOperatorRequest struct {
	UserID         string  `json:"user_id" validate:"required,uuid"`
	OperatorCode   string  `json:"operator_code" validate:"required,max=20"`
	CompanyName    string  `json:"company_name" validate:"required,max=255"`
	LicenseNumber  string  `json:"license_number" validate:"required,max=100"`
	ContactPerson  string  `json:"contact_person" validate:"required,max=255"`
	ContactPhone   string  `json:"contact_phone" validate:"required,max=20"`
	ContactEmail   string  `json:"contact_email" validate:"required,email,max=255"`
	Address        string  `json:"address" validate:"required"`
	City           string  `json:"city" validate:"required,max=100"`
	Province       string  `json:"province" validate:"required,max=100"`
	Country        string  `json:"country" validate:"required,max=100"`
	PostalCode     string  `json:"postal_code" validate:"required,max=20"`
	Website        *string `json:"website" validate:"omitempty,url,max=500"`
	OperatorType   string  `json:"operator_type" validate:"required,max=100"`
	EstablishedAt  *int64  `json:"established_at"`
	LicenseExpiry  *int64  `json:"license_expiry"`
	Notes          *string `json:"notes" validate:"omitempty,max=1000"`
}

type UpdateOperatorRequest struct {
	ID             string  `json:"-" validate:"required,max=100,uuid"`
	OperatorCode   *string `json:"operator_code" validate:"omitempty,max=20"`
	CompanyName    *string `json:"company_name" validate:"omitempty,max=255"`
	LicenseNumber  *string `json:"license_number" validate:"omitempty,max=100"`
	ContactPerson  *string `json:"contact_person" validate:"omitempty,max=255"`
	ContactPhone   *string `json:"contact_phone" validate:"omitempty,max=20"`
	ContactEmail   *string `json:"contact_email" validate:"omitempty,email,max=255"`
	Address        *string `json:"address"`
	City           *string `json:"city" validate:"omitempty,max=100"`
	Province       *string `json:"province" validate:"omitempty,max=100"`
	Country        *string `json:"country" validate:"omitempty,max=100"`
	PostalCode     *string `json:"postal_code" validate:"omitempty,max=20"`
	Website        *string `json:"website" validate:"omitempty,url,max=500"`
	OperatorType   *string `json:"operator_type" validate:"omitempty,max=100"`
	Status         *string `json:"status" validate:"omitempty,oneof=active inactive suspended"`
	EstablishedAt  *int64  `json:"established_at"`
	LicenseExpiry  *int64  `json:"license_expiry"`
	IsActive       *bool   `json:"is_active"`
	Notes          *string `json:"notes" validate:"omitempty,max=1000"`
}

type GetOperatorRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteOperatorRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type ListOperatorRequest struct {
	Page         int     `json:"page" validate:"min=1"`
	Size         int     `json:"size" validate:"min=1,max=100"`
	IsActive     *bool   `json:"is_active"`
	Status       *string `json:"status"`
	OperatorType *string `json:"operator_type"`
	Country      *string `json:"country"`
	Province     *string `json:"province"`
	City         *string `json:"city"`
	CompanyName  *string `json:"company_name"`
}

type GetOperatorByUserRequest struct {
	UserID string `json:"-" validate:"required,max=100,uuid"`
}