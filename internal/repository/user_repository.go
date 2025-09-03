package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) CountByUsernameAndEmail(db *gorm.DB, username, email, excludeID string) (int64, error) {
	var count int64
	query := db.Model(&entity.User{}).Where("username = ? OR email = ?", username, email)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	return count, query.Count(&count).Error
}

func (r *UserRepository) FindByUsername(db *gorm.DB, user *entity.User, username string) error {
	return db.Where("username = ?", username).First(user).Error
}

func (r *UserRepository) CountByUsername(db *gorm.DB, username, excludeID string) (int64, error) {
	var count int64
	query := db.Model(&entity.User{}).Where("username = ?", username)
	if excludeID != "" {
		query = query.Where("id != ?", excludeID)
	}
	return count, query.Count(&count).Error
}
