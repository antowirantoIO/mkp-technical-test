package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	domain "mkp-boarding-test/internal/domain/repository"
	baseRepo "mkp-boarding-test/internal/infrastructure/repository/base"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	baseRepo.Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) domain.UserRepository {
	return &UserRepositoryImpl{
		Log: log,
	}
}

func (r *UserRepositoryImpl) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepositoryImpl) CountByUsernameAndEmail(db *gorm.DB, username, email, excludeID string) (int64, error) {
	var count int64
	query := db.Model(&entity.User{}).Where("username = ? OR email = ?", username, email)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	return count, query.Count(&count).Error
}

func (r *UserRepositoryImpl) FindByUsername(db *gorm.DB, user *entity.User, username string) error {
	return db.Where("username = ?", username).First(user).Error
}

func (r *UserRepositoryImpl) CountByUsername(db *gorm.DB, username, excludeID string) (int64, error) {
	var count int64
	query := db.Model(&entity.User{}).Where("username = ?", username)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	return count, query.Count(&count).Error
}
