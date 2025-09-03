package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	// Base CRUD operations
	Create(db *gorm.DB, role *entity.Role) error
	Update(db *gorm.DB, role *entity.Role) error
	Delete(db *gorm.DB, role *entity.Role) error
	FindById(db *gorm.DB, role *entity.Role, id any) error
	CountById(db *gorm.DB, id any) (int64, error)

	// Custom operations
	FindByName(db *gorm.DB, role *entity.Role, name string) error
	FindAllActive(db *gorm.DB) ([]entity.Role, error)
	CountByName(db *gorm.DB, name string, excludeID string) (int64, error)
}