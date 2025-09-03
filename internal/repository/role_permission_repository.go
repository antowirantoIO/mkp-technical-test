package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RolePermissionRepository struct {
	Repository[entity.RolePermission]
	Log *logrus.Logger
}

func NewRolePermissionRepository(log *logrus.Logger) *RolePermissionRepository {
	return &RolePermissionRepository{
		Log: log,
	}
}

func (r *RolePermissionRepository) FindByRoleIDAndPermissionID(db *gorm.DB, rolePermission *entity.RolePermission, roleID string, permissionID string) error {
	return db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).First(rolePermission).Error
}

func (r *RolePermissionRepository) FindAllByRoleID(db *gorm.DB, roleID string) ([]entity.RolePermission, error) {
	var rolePermissions []entity.RolePermission
	if err := db.Preload("Permission").Where("role_id = ?", roleID).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (r *RolePermissionRepository) FindAllByPermissionID(db *gorm.DB, permissionID string) ([]entity.RolePermission, error) {
	var rolePermissions []entity.RolePermission
	if err := db.Preload("Role").Where("permission_id = ?", permissionID).Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}

func (r *RolePermissionRepository) DeleteByRoleIDAndPermissionID(db *gorm.DB, roleID string, permissionID string) error {
	return db.Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&entity.RolePermission{}).Error
}

func (r *RolePermissionRepository) DeleteAllByRoleID(db *gorm.DB, roleID string) error {
	return db.Where("role_id = ?", roleID).Delete(&entity.RolePermission{}).Error
}