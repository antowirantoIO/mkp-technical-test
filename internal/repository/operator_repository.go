package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OperatorRepository struct {
	Repository[entity.Operator]
	Log *logrus.Logger
}

func NewOperatorRepository(log *logrus.Logger) *OperatorRepository {
	return &OperatorRepository{
		Log: log,
	}
}

func (r *OperatorRepository) FindByUserID(db *gorm.DB, operator *entity.Operator, userID string) error {
	return db.Where("user_id = ? AND deleted_at IS NULL", userID).First(operator).Error
}

func (r *OperatorRepository) FindByOperatorCode(db *gorm.DB, operator *entity.Operator, operatorCode string) error {
	return db.Where("operator_code = ? AND deleted_at IS NULL", operatorCode).First(operator).Error
}

func (r *OperatorRepository) FindByLicenseNumber(db *gorm.DB, operator *entity.Operator, licenseNumber string) error {
	return db.Where("license_number = ? AND deleted_at IS NULL", licenseNumber).First(operator).Error
}

func (r *OperatorRepository) FindAllActive(db *gorm.DB) ([]entity.Operator, error) {
	var operators []entity.Operator
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&operators).Error; err != nil {
		return nil, err
	}
	return operators, nil
}

func (r *OperatorRepository) CountByOperatorCode(db *gorm.DB, operatorCode string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Operator{}).Where("operator_code = ? AND deleted_at IS NULL", operatorCode)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *OperatorRepository) CountByLicenseNumber(db *gorm.DB, licenseNumber string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Operator{}).Where("license_number = ? AND deleted_at IS NULL", licenseNumber)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}