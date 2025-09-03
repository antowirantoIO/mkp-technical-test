package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	"gorm.io/gorm"
)

type OperatorRepository interface {
	// Base CRUD operations
	Create(db *gorm.DB, operator *entity.Operator) error
	Update(db *gorm.DB, operator *entity.Operator) error
	Delete(db *gorm.DB, operator *entity.Operator) error
	FindById(db *gorm.DB, operator *entity.Operator, id any) error
	CountById(db *gorm.DB, id any) (int64, error)

	// Custom operations
	FindByUserID(db *gorm.DB, operator *entity.Operator, userID string) error
	FindByOperatorCode(db *gorm.DB, operator *entity.Operator, operatorCode string) error
	FindByLicenseNumber(db *gorm.DB, operator *entity.Operator, licenseNumber string) error
	FindAllActive(db *gorm.DB) ([]entity.Operator, error)
	CountByOperatorCode(db *gorm.DB, operatorCode string, excludeID string) (int64, error)
	CountByLicenseNumber(db *gorm.DB, licenseNumber string, excludeID string) (int64, error)
}