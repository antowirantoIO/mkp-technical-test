package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	baseRepo "mkp-boarding-test/internal/infrastructure/repository/base"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRoleRepositoryImpl struct {
	baseRepo.Repository[entity.UserRole]
	Log *logrus.Logger
}

func NewUserRoleRepository(log *logrus.Logger) *UserRoleRepositoryImpl {
	return &UserRoleRepositoryImpl{
		Log: log,
	}
}

func (r *UserRoleRepositoryImpl) FindByUserIDAndRoleID(db *gorm.DB, userRole *entity.UserRole, userID string, roleID string) error {
	return db.Where("user_id = ? AND role_id = ?", userID, roleID).First(userRole).Error
}

func (r *UserRoleRepositoryImpl) FindAllByUserID(db *gorm.DB, userID string) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole
	if err := db.Preload("Role").Where("user_id = ?", userID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepositoryImpl) FindAllByRoleID(db *gorm.DB, roleID string) ([]entity.UserRole, error) {
	var userRoles []entity.UserRole
	if err := db.Preload("User").Where("role_id = ?", roleID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepositoryImpl) DeleteByUserIDAndRoleID(db *gorm.DB, userID string, roleID string) error {
	return db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&entity.UserRole{}).Error
}

func (r *UserRoleRepositoryImpl) DeleteAllByUserID(db *gorm.DB, userID string) error {
	return db.Where("user_id = ?", userID).Delete(&entity.UserRole{}).Error
}
