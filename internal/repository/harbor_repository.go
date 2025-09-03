package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HarborRepository struct {
	Repository[entity.Harbor]
	Log *logrus.Logger
}

func NewHarborRepository(log *logrus.Logger) *HarborRepository {
	return &HarborRepository{
		Log: log,
	}
}

func (r *HarborRepository) FindByHarborCode(db *gorm.DB, harbor *entity.Harbor, harborCode string) error {
	return db.Where("harbor_code = ? AND deleted_at IS NULL", harborCode).First(harbor).Error
}

func (r *HarborRepository) FindByUNLocode(db *gorm.DB, harbor *entity.Harbor, unLocode string) error {
	return db.Where("un_locode = ? AND deleted_at IS NULL", unLocode).First(harbor).Error
}

func (r *HarborRepository) FindByCountry(db *gorm.DB, country string) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("country = ? AND deleted_at IS NULL", country).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepository) FindByProvince(db *gorm.DB, province string) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("province = ? AND deleted_at IS NULL", province).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepository) FindByCity(db *gorm.DB, city string) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("city = ? AND deleted_at IS NULL", city).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepository) FindAllActive(db *gorm.DB) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepository) CountByHarborCode(db *gorm.DB, harborCode string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Harbor{}).Where("harbor_code = ? AND deleted_at IS NULL", harborCode)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *HarborRepository) CountByUNLocode(db *gorm.DB, unLocode string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Harbor{}).Where("un_locode = ? AND deleted_at IS NULL", unLocode)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}