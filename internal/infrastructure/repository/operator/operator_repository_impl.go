package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	domain "mkp-boarding-test/internal/domain/repository"
	baseRepo "mkp-boarding-test/internal/infrastructure/repository/base"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OperatorRepositoryImpl struct {
	baseRepo.Repository[entity.Operator]
	Log *logrus.Logger
}

var _ domain.OperatorRepository = (*OperatorRepositoryImpl)(nil)

func NewOperatorRepository(log *logrus.Logger) *OperatorRepositoryImpl {
	return &OperatorRepositoryImpl{
		Log: log,
	}
}

func (r *OperatorRepositoryImpl) FindByUserID(db *gorm.DB, operator *entity.Operator, userID string) error {
	return db.Where("user_id = ? AND deleted_at IS NULL", userID).First(operator).Error
}

func (r *OperatorRepositoryImpl) FindByOperatorCode(db *gorm.DB, operator *entity.Operator, operatorCode string) error {
	return db.Where("operator_code = ? AND deleted_at IS NULL", operatorCode).First(operator).Error
}

func (r *OperatorRepositoryImpl) FindByLicenseNumber(db *gorm.DB, operator *entity.Operator, licenseNumber string) error {
	return db.Where("license_number = ? AND deleted_at IS NULL", licenseNumber).First(operator).Error
}

func (r *OperatorRepositoryImpl) FindAllActive(db *gorm.DB) ([]entity.Operator, error) {
	var operators []entity.Operator
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&operators).Error; err != nil {
		return nil, err
	}
	return operators, nil
}

func (r *OperatorRepositoryImpl) CountByOperatorCode(db *gorm.DB, operatorCode string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Operator{}).Where("operator_code = ? AND deleted_at IS NULL", operatorCode)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *OperatorRepositoryImpl) CountByLicenseNumber(db *gorm.DB, licenseNumber string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Operator{}).Where("license_number = ? AND deleted_at IS NULL", licenseNumber)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}
