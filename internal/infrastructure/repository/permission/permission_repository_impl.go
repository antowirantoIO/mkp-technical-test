package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	domain "mkp-boarding-test/internal/domain/repository"
	baseRepo "mkp-boarding-test/internal/infrastructure/repository/base"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PermissionRepositoryImpl struct {
	baseRepo.Repository[entity.Permission]
	Log *logrus.Logger
}

func NewPermissionRepository(log *logrus.Logger) domain.PermissionRepository {
	return &PermissionRepositoryImpl{
		Log: log,
	}
}

func (r *PermissionRepositoryImpl) FindByName(db *gorm.DB, permission *entity.Permission, name string) error {
	return db.Where("name = ? AND deleted_at IS NULL", name).First(permission).Error
}

func (r *PermissionRepositoryImpl) FindByResourceAndAction(db *gorm.DB, permission *entity.Permission, resource string, action string) error {
	return db.Where("resource = ? AND action = ? AND deleted_at IS NULL", resource, action).First(permission).Error
}

func (r *PermissionRepositoryImpl) FindAllActive(db *gorm.DB) ([]entity.Permission, error) {
	var permissions []entity.Permission
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *PermissionRepositoryImpl) CountByName(db *gorm.DB, name string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Permission{}).Where("name = ? AND deleted_at IS NULL", name)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}
