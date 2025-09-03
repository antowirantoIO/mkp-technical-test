package repository

import (
	"mkp-boarding-test/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Base CRUD operations
	Create(db *gorm.DB, user *entity.User) error
	Update(db *gorm.DB, user *entity.User) error
	Delete(db *gorm.DB, user *entity.User) error
	FindById(db *gorm.DB, user *entity.User, id any) error
	CountById(db *gorm.DB, id any) (int64, error)

	// Custom operations
	FindByToken(db *gorm.DB, user *entity.User, token string) error
	CountByUsernameAndEmail(db *gorm.DB, username, email, excludeID string) (int64, error)
	FindByUsername(db *gorm.DB, user *entity.User, username string) error
	CountByUsername(db *gorm.DB, username, excludeID string) (int64, error)
}