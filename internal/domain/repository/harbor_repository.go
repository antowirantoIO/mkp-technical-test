package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	"gorm.io/gorm"
)

type HarborRepository interface {
	// Base CRUD operations
	Create(db *gorm.DB, harbor *entity.Harbor) error
	Update(db *gorm.DB, harbor *entity.Harbor) error
	Delete(db *gorm.DB, harbor *entity.Harbor) error
	FindById(db *gorm.DB, harbor *entity.Harbor, id any) error
	CountById(db *gorm.DB, id any) (int64, error)

	// Custom operations
	FindByHarborCode(db *gorm.DB, harbor *entity.Harbor, harborCode string) error
	FindByUNLocode(db *gorm.DB, harbor *entity.Harbor, unLocode string) error
	FindByCountry(db *gorm.DB, country string) ([]entity.Harbor, error)
	FindByProvince(db *gorm.DB, province string) ([]entity.Harbor, error)
	FindByCity(db *gorm.DB, city string) ([]entity.Harbor, error)
	FindAllActive(db *gorm.DB) ([]entity.Harbor, error)
	CountByHarborCode(db *gorm.DB, harborCode string, excludeID string) (int64, error)
	CountByUNLocode(db *gorm.DB, unLocode string, excludeID string) (int64, error)
}