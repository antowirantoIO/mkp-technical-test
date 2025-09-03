package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	Repository[entity.UserRole]
	Log *logrus.Logger
}

func NewUserRoleRepository(log *logrus.Logger) *UserRoleRepository {
	return &UserRoleRepository{
		Log: log,
	}
}

func (r *UserRoleRepository) FindByUserIDAndRoleID(db *gorm.DB, userRole *entity.UserRole, userID string, roleID string) error {
	return db.Where("user_id = ? AND role_id = ?", userID, roleID).First(userRole).Error
}

func (r *UserRoleRepository) FindAllByUserID(db *gorm.DB, userID string) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole
	if err := db.Preload("Role").Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepository) FindAllByRoleID(db *gorm.DB, roleID string) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole
	if err := db.Preload("User").Where("role_id = ?", roleID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepository) DeleteByUserIDAndRoleID(db *gorm.DB, userID string, roleID string) error {
	return db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&entity.UserRole{}).Error
}

func (r *UserRoleRepository) DeleteAllByUserID(db *gorm.DB, userID string) error {
	return db.Where("user_id = ?", userID).Delete(&entity.UserRole{}).Error
}