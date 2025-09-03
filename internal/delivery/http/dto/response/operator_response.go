package response

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