package repository

import (
	"mkp-boarding-test/internal/domain/entity"

	"gorm.io/gorm"
)

type ShipRepository interface {
	// Base CRUD operations
	Create(db *gorm.DB, ship *entity.Ship) error
	Update(db *gorm.DB, ship *entity.Ship) error
	Delete(db *gorm.DB, ship *entity.Ship) error
	FindById(db *gorm.DB, ship *entity.Ship, id any) error
	CountById(db *gorm.DB, id any) (int64, error)

	// Custom operations
	FindByOperatorID(db *gorm.DB, operatorID string) ([]entity.Ship, error)
	FindByIMONumber(db *gorm.DB, ship *entity.Ship, imoNumber string) error
	FindByCallSign(db *gorm.DB, ship *entity.Ship, callSign string) error
	FindByMMSI(db *gorm.DB, ship *entity.Ship, mmsi string) error
	FindAllActive(db *gorm.DB) ([]entity.Ship, error)
	CountByIMONumber(db *gorm.DB, imoNumber string, excludeID string) (int64, error)
	CountByCallSign(db *gorm.DB, callSign string, excludeID string) (int64, error)
	CountByMMSI(db *gorm.DB, mmsi string, excludeID string) (int64, error)
	CountByShipNameAndOperatorID(db *gorm.DB, shipName string, operatorID string, excludeID string) (int64, error)
}
