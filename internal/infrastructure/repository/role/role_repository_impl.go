package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	domain "mkp-boarding-test/internal/domain/repository"
	baseRepo "mkp-boarding-test/internal/infrastructure/repository/base"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	baseRepo.Repository[entity.Role]
	Log *logrus.Logger
}

var _ domain.RoleRepository = (*RoleRepositoryImpl)(nil)

func NewRoleRepository(log *logrus.Logger) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{
		Log: log,
	}
}

func (r *RoleRepositoryImpl) FindByName(db *gorm.DB, role *entity.Role, name string) error {
	return db.Where("name = ? AND deleted_at IS NULL", name).First(role).Error
}

func (r *RoleRepositoryImpl) FindAllActive(db *gorm.DB) ([]entity.Role, error) {
	var roles []entity.Role
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) CountByName(db *gorm.DB, name string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Role{}).Where("name = ? AND deleted_at IS NULL", name)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}
