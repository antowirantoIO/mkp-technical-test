package entity

// UserRole is a struct that represents a user_role junction entity
type UserRole struct {
	UserID    string `gorm:"column:user_id;primaryKey"`
	RoleID    string `gorm:"column:role_id;primaryKey"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`

	// Relations
	User User `gorm:"foreignKey:user_id;references:id"`
	Role Role `gorm:"foreignKey:role_id;references:id"`
}

func (ur *UserRole) TableName() string {
	return "user_roles"
}