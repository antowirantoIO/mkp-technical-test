package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	// Base CRUD operations
	Create(db *gorm.DB, permission *entity.Permission) error
	Update(db *gorm.DB, permission *entity.Permission) error
	Delete(db *gorm.DB, permission *entity.Permission) error
	FindById(db *gorm.DB, permission *entity.Permission, id any) error
	CountById(db *gorm.DB, id any) (int64, error)

	// Custom operations
	FindByName(db *gorm.DB, permission *entity.Permission, name string) error
	FindByResourceAndAction(db *gorm.DB, permission *entity.Permission, resource string, action string) error
	FindAllActive(db *gorm.DB) ([]entity.Permission, error)
	CountByName(db *gorm.DB, name string, excludeID string) (int64, error)
}