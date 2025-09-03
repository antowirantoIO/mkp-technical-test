package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Repository[entity.Role]
	Log *logrus.Logger
}

func NewRoleRepository(log *logrus.Logger) *RoleRepository {
	return &RoleRepository{
		Log: log,
	}
}

func (r *RoleRepository) FindByName(db *gorm.DB, role *entity.Role, name string) error {
	return db.Where("name = ? AND deleted_at IS NULL", name).First(role).Error
}

func (r *RoleRepository) FindAllActive(db *gorm.DB) ([]entity.Role, error) {
	var roles []entity.Role
	if err := db.Where("is_active = ? AND deleted_at IS NULL", true).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) CountByName(db *gorm.DB, name string, excludeID string) (int64, error) {
	var total int64
	query := db.Model(&entity.Role{}).Where("name = ? AND deleted_at IS NULL", name)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	err := query.Count(&total).Error
	return total, err
}