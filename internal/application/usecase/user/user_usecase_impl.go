package user

import (
	"context"
	"time"

	"mkp-boarding-test/internal/delivery/http/dto/request"
	"mkp-boarding-test/internal/delivery/http/dto/response"
	"mkp-boarding-test/internal/domain/entity"
	"mkp-boarding-test/internal/domain/repository"
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/gateway/messaging"
	"mkp-boarding-test/internal/model/converter"
	"mkp-boarding-test/pkg/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCaseImpl struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository repository.UserRepository
	UserProducer   *messaging.UserProducer
	JWTService     service.JWTService
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	userRepository repository.UserRepository, userProducer *messaging.UserProducer, jwtService service.JWTService) usecase.UserUseCase {
	return &UserUseCaseImpl{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
		UserProducer:   userProducer,
		JWTService:     jwtService,
	}
}

func (c *UserUseCaseImpl) Verify(ctx context.Context, request *request.VerifyUserRequest) (*response.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByToken(tx, user, request.Token); err != nil {
		c.Log.WithError(err).Error("failed to find user by token")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCaseImpl) Create(ctx context.Context, request *request.RegisterUserRequest) (*response.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	if count, err := c.UserRepository.CountByUsernameAndEmail(tx, request.Username, request.Email, ""); err != nil {
		c.Log.WithError(err).Error("failed to count user by username and email")
		return nil, fiber.ErrInternalServerError
	} else if count > 0 {
		c.Log.Error("user already exists")
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.WithError(err).Error("failed to generate bcrypt hash")
		return nil, fiber.ErrInternalServerError
	}

	firstName := ""
	if request.FirstName != nil {
		firstName = *request.FirstName
	}
	lastName := ""
	if request.LastName != nil {
		lastName = *request.LastName
	}

	user := &entity.User{
		ID:         uuid.NewString(),
		Username:   request.Username,
		Email:      request.Email,
		Password:   string(password),
		FirstName:  firstName,
		LastName:   lastName,
		Phone:      request.Phone,
		IsActive:   true,
		IsVerified: false,
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.WithError(err).Error("failed to create user")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCaseImpl) Login(ctx context.Context, request *request.LoginUserRequest) (*response.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByUsername(tx, user, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find user by username")
		return nil, fiber.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Log.WithError(err).Error("invalid password")
		return nil, fiber.ErrUnauthorized
	}

	// Generate JWT tokens
	token, err := c.JWTService.GenerateToken(user)
	if err != nil {
		c.Log.WithError(err).Error("failed to generate JWT token")
		return nil, fiber.ErrInternalServerError
	}

	refreshToken, err := c.JWTService.GenerateRefreshToken(user)
	if err != nil {
		c.Log.WithError(err).Error("failed to generate refresh token")
		return nil, fiber.ErrInternalServerError
	}

	// Update user with tokens
	now := time.Now()
	tokenExpiresAt := now.Add(24 * time.Hour).Unix()       // Token expires in 24 hours
	refreshExpiresAt := now.Add(7 * 24 * time.Hour).Unix() // Refresh token expires in 7 days
	user.Token = &token
	user.TokenExpiresAt = &tokenExpiresAt
	user.RefreshToken = &refreshToken
	user.RefreshExpiresAt = &refreshExpiresAt

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.WithError(err).Error("failed to update user with tokens")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	response := converter.UserToResponse(user)
	response.Token = token
	response.RefreshToken = refreshToken

	return response, nil
}

func (c *UserUseCaseImpl) Current(ctx context.Context, request *request.GetUserRequest) (*response.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindById(tx, user, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find user by id")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCaseImpl) Logout(ctx context.Context, request *request.LogoutUserRequest) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return false, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindById(tx, user, request.ID); err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return false, fiber.ErrNotFound
	}

	// Clear tokens from database
	user.Token = nil
	user.TokenExpiresAt = nil
	user.RefreshToken = nil
	user.RefreshExpiresAt = nil

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.WithError(err).Error("failed to clear user tokens")
		return false, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func (c *UserUseCaseImpl) Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	if err := c.UserRepository.FindById(tx, user, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find user by id")
		return nil, fiber.ErrNotFound
	}

	if request.FirstName != nil {
		user.FirstName = *request.FirstName
	}

	if request.LastName != nil {
		user.LastName = *request.LastName
	}

	if request.Phone != nil {
		user.Phone = request.Phone
	}

	if request.Password != nil && *request.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(*request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Log.WithError(err).Error("failed to generate bcrypt hash")
			return nil, fiber.ErrInternalServerError
		}
		user.Password = string(password)
	}

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.WithError(err).Error("failed to update user")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCaseImpl) FindByRoleID(ctx context.Context, roleID string) ([]*response.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if roleID == "" {
		c.Log.Error("roleID is required")
		return nil, fiber.ErrBadRequest
	}

	users, err := c.UserRepository.FindByRoleID(tx, roleID)
	if err != nil {
		c.Log.WithError(err).Error("failed to find users by role id")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponseList(users), nil
}
