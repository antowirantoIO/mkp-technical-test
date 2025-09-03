package entity

// User is a struct that represents a user entity
type User struct {
	ID                string     `gorm:"column:id;primaryKey"`
	Username          string     `gorm:"column:username;uniqueIndex"`
	Email             string     `gorm:"column:email;uniqueIndex"`
	Password          string     `gorm:"column:password"`
	FirstName         string     `gorm:"column:first_name"`
	LastName          string     `gorm:"column:last_name"`
	Phone             *string    `gorm:"column:phone"`
	Avatar            *string    `gorm:"column:avatar"`
	IsActive          bool       `gorm:"column:is_active;default:true"`
	IsVerified        bool       `gorm:"column:is_verified;default:false"`
	LastLoginAt       *int64     `gorm:"column:last_login_at"`
	EmailVerifiedAt   *int64     `gorm:"column:email_verified_at"`
	PasswordChangedAt *int64     `gorm:"column:password_changed_at"`
	CreatedAt         int64      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt         int64      `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt         *int64     `gorm:"column:deleted_at"`
	
	// Relations
	Contacts  []Contact  `gorm:"foreignKey:user_id;references:id"`
	Operator  *Operator  `gorm:"foreignKey:user_id;references:id"`
	UserRoles []UserRole `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
