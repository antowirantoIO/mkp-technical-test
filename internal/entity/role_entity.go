package entity

// Role is a struct that represents a role entity
type Role struct {
	ID          string `gorm:"column:id;primaryKey"`
	Name        string `gorm:"column:name;uniqueIndex"`
	DisplayName string `gorm:"column:display_name"`
	Description *string `gorm:"column:description"`
	IsActive    bool   `gorm:"column:is_active;default:true"`
	IsSystem    bool   `gorm:"column:is_system;default:false"`
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   *int64 `gorm:"column:deleted_at"`

	// Relations
	UserRoles       []UserRole       `gorm:"foreignKey:role_id;references:id"`
	RolePermissions []RolePermission `gorm:"foreignKey:role_id;references:id"`
}

func (r *Role) TableName() string {
	return "roles"
}