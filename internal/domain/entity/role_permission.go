package entity

// RolePermission is a struct that represents a role_permission junction entity
type RolePermission struct {
	RoleID       string `gorm:"column:role_id;primaryKey"`
	PermissionID string `gorm:"column:permission_id;primaryKey"`
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime:milli"`

	// Relations
	Role       Role       `gorm:"foreignKey:role_id;references:id"`
	Permission Permission `gorm:"foreignKey:permission_id;references:id"`
}

func (rp *RolePermission) TableName() string {
	return "role_permissions"
}