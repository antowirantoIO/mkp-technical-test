package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShipRepository struct {
	Repository[entity.Ship]
	Log *logrus.Logger
}

func NewShipRepository(log *logrus.Logger) *ShipRepository {
	return &ShipRepository{
		Log: log,
	}
}

func (r *ShipRepository) FindByOperatorID(db *gorm.DB, operatorID string) ([]entity.Ship, error) {
	var ships []entity.Ship
	if err := db.Where("operator_id = ? AND deleted_at IS NULL", operatorID).Find(&ships).Error; err != nil {
		return nil, err
	}
	return ships, nil
}

func (r *ShipRepository) FindByIMONumber(db *gorm.DB, ship *entity.Ship, imoNumber string) error {
	return db.Where("imo_number = ? AND deleted_at IS NULL", imoNumber).First(ship).Error
}

func (r *ShipRepository) FindByCallSign(db *gorm.DB, ship *entity.Ship, callSign string) error {
	return db.Where("call_sign = ? AND deleted_at IS NULL", callSign).First(ship).Error
}

func (r *ShipRepository) FindByMMSI(db *gorm.DB, ship *entity.Ship, mmsi string) error {
	return db.Where("mmsi = ? AND deleted_at IS NULL", mmsi).First(ship).Error
}

func (r *ShipRepository) FindAllActive(db *gorm.DB) ([]entity.Ship, error) {
	var ships []entity.Ship
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&ships).Error; err != nil {
		return nil, err
	}
	return ships, nil
}

func (r *ShipRepository) CountByIMONumber(db *gorm.DB, imoNumber string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Ship{}).Where("imo_number = ? AND deleted_at IS NULL", imoNumber)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *ShipRepository) CountByCallSign(db *gorm.DB, callSign string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Ship{}).Where("call_sign = ? AND deleted_at IS NULL", callSign)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *ShipRepository) CountByMMSI(db *gorm.DB, mmsi string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Ship{}).Where("mmsi = ? AND deleted_at IS NULL", mmsi)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *ShipRepository) CountByShipNameAndOperatorID(db *gorm.DB, shipName string, operatorID string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Ship{}).Where("ship_name = ? AND operator_id = ? AND deleted_at IS NULL", shipName, operatorID)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}