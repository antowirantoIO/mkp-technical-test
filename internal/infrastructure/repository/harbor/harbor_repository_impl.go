package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	domain "mkp-boarding-test/internal/domain/repository"
	baseRepo "mkp-boarding-test/internal/infrastructure/repository/base"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HarborRepositoryImpl struct {
	baseRepo.Repository[entity.Harbor]
	Log *logrus.Logger
}

var _ domain.HarborRepository = (*HarborRepositoryImpl)(nil)

func NewHarborRepository(log *logrus.Logger) domain.HarborRepository {
	return &HarborRepositoryImpl{
		Log: log,
	}
}

func (r *HarborRepositoryImpl) FindByHarborCode(db *gorm.DB, harbor *entity.Harbor, harborCode string) error {
	return db.Where("harbor_code = ? AND deleted_at IS NULL", harborCode).First(harbor).Error
}

func (r *HarborRepositoryImpl) FindByUNLocode(db *gorm.DB, harbor *entity.Harbor, unLocode string) error {
	return db.Where("un_locode = ? AND deleted_at IS NULL", unLocode).First(harbor).Error
}

func (r *HarborRepositoryImpl) FindByCountry(db *gorm.DB, country string) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("country = ? AND deleted_at IS NULL", country).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepositoryImpl) FindByProvince(db *gorm.DB, province string) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("province = ? AND deleted_at IS NULL", province).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepositoryImpl) FindByCity(db *gorm.DB, city string) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("city = ? AND deleted_at IS NULL", city).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepositoryImpl) FindAllActive(db *gorm.DB) ([]entity.Harbor, error) {
	var harbors []entity.Harbor
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&harbors).Error; err != nil {
		return nil, err
	}
	return harbors, nil
}

func (r *HarborRepositoryImpl) CountByHarborCode(db *gorm.DB, harborCode string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Harbor{}).Where("harbor_code = ? AND deleted_at IS NULL", harborCode)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}

func (r *HarborRepositoryImpl) CountByUNLocode(db *gorm.DB, unLocode string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Harbor{}).Where("un_locode = ? AND deleted_at IS NULL", unLocode)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}
