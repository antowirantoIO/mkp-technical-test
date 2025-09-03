package entity

// Operator is a struct that represents an operator entity
type Operator struct {
	ID             string  `gorm:"column:id;primaryKey"`
	UserID         string  `gorm:"column:user_id;uniqueIndex"`
	OperatorCode   string  `gorm:"column:operator_code;uniqueIndex"`
	CompanyName    string  `gorm:"column:company_name"`
	LicenseNumber  string  `gorm:"column:license_number;uniqueIndex"`
	ContactPerson  string  `gorm:"column:contact_person"`
	ContactPhone   string  `gorm:"column:contact_phone"`
	ContactEmail   string  `gorm:"column:contact_email"`
	Address        string  `gorm:"column:address"`
	City           string  `gorm:"column:city"`
	Province       string  `gorm:"column:province"`
	Country        string  `gorm:"column:country"`
	PostalCode     string  `gorm:"column:postal_code"`
	Website        *string `gorm:"column:website"`
	OperatorType   string  `gorm:"column:operator_type"`
	Status         string  `gorm:"column:status;default:active"`
	EstablishedAt  *int64  `gorm:"column:established_at"`
	LicenseExpiry  *int64  `gorm:"column:license_expiry"`
	IsActive       bool    `gorm:"column:is_active;default:true"`
	Notes          *string `gorm:"column:notes"`
	CreatedAt      int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt      int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt      *int64  `gorm:"column:deleted_at"`

	// Relations
	User  User   `gorm:"foreignKey:user_id;references:id"`
	Ships []Ship `gorm:"foreignKey:operator_id;references:id"`
}

func (o *Operator) TableName() string {
	return "operators"
}