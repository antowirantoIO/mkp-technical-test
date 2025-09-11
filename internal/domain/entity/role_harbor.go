package entity

// RolePermission is a struct that represents a role_permission junction entity
type RoleHarbor struct {
	RoleID   string `gorm:"column:role_id;primaryKey"`
	HarborId string `gorm:"column:harbor_id;primaryKey"`

	// Relations
	Role   Role   `gorm:"foreignKey:role_id;references:id"`
	Harbor Harbor `gorm:"foreignKey:harbor_id;references:id"`
}

func (rp *RoleHarbor) TableName() string {
	return "role_harbors"
}
