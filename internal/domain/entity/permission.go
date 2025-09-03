package entity

// Permission is a struct that represents a permission entity
type Permission struct {
	ID          string  `gorm:"column:id;primaryKey"`
	Name        string  `gorm:"column:name;uniqueIndex"`
	DisplayName string  `gorm:"column:display_name"`
	Description *string `gorm:"column:description"`
	Resource    string  `gorm:"column:resource"`
	Action      string  `gorm:"column:action"`
	IsActive    bool    `gorm:"column:is_active;default:true"`
	IsSystem    bool    `gorm:"column:is_system;default:false"`
	CreatedAt   int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt   *int64  `gorm:"column:deleted_at"`

	// Relations
	RolePermissions []RolePermission `gorm:"foreignKey:permission_id;references:id"`
}

func (p *Permission) TableName() string {
	return "permissions"
}