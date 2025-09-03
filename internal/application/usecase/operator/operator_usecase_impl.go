package operator

import (
	"context"
	"mkp-boarding-test/internal/domain/entity"
	"mkp-boarding-test/internal/domain/repository"
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/internal/model/converter"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OperatorUseCaseImpl struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	OperatorRepository repository.OperatorRepository
}

func NewOperatorUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	operatorRepository repository.OperatorRepository) usecase.OperatorUseCase {
	return &OperatorUseCaseImpl{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		OperatorRepository: operatorRepository,
	}
}

func (c *OperatorUseCaseImpl) Create(ctx context.Context, request *model.CreateOperatorRequest) (*model.OperatorResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	// Check if operator code already exists
	if count, err := c.OperatorRepository.CountByOperatorCode(tx, request.OperatorCode, ""); err != nil {
		c.Log.WithError(err).Error("failed to count operator by code")
		return nil, fiber.ErrInternalServerError
	} else if count > 0 {
		c.Log.Error("operator code already exists")
		return nil, fiber.ErrConflict
	}

	operator := &entity.Operator{
		ID:            uuid.NewString(),
		UserID:        request.UserID,
		OperatorCode:  request.OperatorCode,
		CompanyName:   request.CompanyName,
		LicenseNumber: request.LicenseNumber,
		ContactPerson: request.ContactPerson,
		ContactPhone:  request.ContactPhone,
		ContactEmail:  request.ContactEmail,
		Address:       request.Address,
		City:          request.City,
		Province:      request.Province,
		Country:       request.Country,
		PostalCode:    request.PostalCode,
		Website:       request.Website,
		OperatorType:  request.OperatorType,
		Status:        "active",
		EstablishedAt: request.EstablishedAt,
		LicenseExpiry: request.LicenseExpiry,
		IsActive:      true,
		Notes:         request.Notes,
	}

	if err := c.OperatorRepository.Create(tx, operator); err != nil {
		c.Log.WithError(err).Error("failed to create operator")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OperatorToResponse(operator), nil
}

func (c *OperatorUseCaseImpl) Update(ctx context.Context, request *model.UpdateOperatorRequest) (*model.OperatorResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	operator := new(entity.Operator)
	if err := c.OperatorRepository.FindById(tx, operator, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find operator")
		return nil, fiber.ErrNotFound
	}

	// Check if new code already exists (if code is being changed)
	if request.OperatorCode != nil && *request.OperatorCode != operator.OperatorCode {
		if count, err := c.OperatorRepository.CountByOperatorCode(tx, *request.OperatorCode, operator.ID); err != nil {
			c.Log.WithError(err).Error("failed to count operator by code")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("operator code already exists")
			return nil, fiber.ErrConflict
		}
		operator.OperatorCode = *request.OperatorCode
	}

	if request.CompanyName != nil {
		operator.CompanyName = *request.CompanyName
	}
	if request.LicenseNumber != nil {
		operator.LicenseNumber = *request.LicenseNumber
	}
	if request.ContactPerson != nil {
		operator.ContactPerson = *request.ContactPerson
	}
	if request.ContactPhone != nil {
		operator.ContactPhone = *request.ContactPhone
	}
	if request.ContactEmail != nil {
		operator.ContactEmail = *request.ContactEmail
	}
	if request.Address != nil {
		operator.Address = *request.Address
	}
	if request.City != nil {
		operator.City = *request.City
	}
	if request.Province != nil {
		operator.Province = *request.Province
	}
	if request.Country != nil {
		operator.Country = *request.Country
	}
	if request.PostalCode != nil {
		operator.PostalCode = *request.PostalCode
	}
	if request.Website != nil {
		operator.Website = request.Website
	}
	if request.OperatorType != nil {
		operator.OperatorType = *request.OperatorType
	}
	if request.Status != nil {
		operator.Status = *request.Status
	}
	if request.EstablishedAt != nil {
		operator.EstablishedAt = request.EstablishedAt
	}
	if request.LicenseExpiry != nil {
		operator.LicenseExpiry = request.LicenseExpiry
	}
	if request.IsActive != nil {
		operator.IsActive = *request.IsActive
	}
	if request.Notes != nil {
		operator.Notes = request.Notes
	}

	if err := c.OperatorRepository.Update(tx, operator); err != nil {
		c.Log.WithError(err).Error("failed to update operator")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OperatorToResponse(operator), nil
}

func (c *OperatorUseCaseImpl) Get(ctx context.Context, request *model.GetOperatorRequest) (*model.OperatorResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	operator := new(entity.Operator)
	if err := c.OperatorRepository.FindById(tx, operator, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find operator")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.OperatorToResponse(operator), nil
}

func (c *OperatorUseCaseImpl) Delete(ctx context.Context, request *model.DeleteOperatorRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	operator := new(entity.Operator)
	if err := c.OperatorRepository.FindById(tx, operator, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find operator")
		return fiber.ErrNotFound
	}

	if err := c.OperatorRepository.Delete(tx, operator); err != nil {
		c.Log.WithError(err).Error("failed to delete operator")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *OperatorUseCaseImpl) List(ctx context.Context, request *model.ListOperatorRequest) (*model.WebResponse[[]model.OperatorResponse], error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	query := tx.Model(&entity.Operator{}).Where("deleted_at IS NULL")

	if request.IsActive != nil {
		query = query.Where("is_active = ?", *request.IsActive)
	}
	if request.CompanyName != nil && *request.CompanyName != "" {
		query = query.Where("company_name ILIKE ?", "%"+*request.CompanyName+"%")
	}
	if request.Country != nil && *request.Country != "" {
		query = query.Where("country = ?", *request.Country)
	}

	// Count total records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.Log.WithError(err).Error("failed to count operators")
		return nil, fiber.ErrInternalServerError
	}

	// Apply pagination
	offset := (request.Page - 1) * request.Size
	query = query.Offset(offset).Limit(request.Size)

	var operators []entity.Operator
	if err := query.Find(&operators).Error; err != nil {
		c.Log.WithError(err).Error("failed to find operators")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.OperatorResponse, len(operators))
	for i, operator := range operators {
		responses[i] = *converter.OperatorToResponse(&operator)
	}

	lastPage := (total + int64(request.Size) - 1) / int64(request.Size)
	if lastPage == 0 {
		lastPage = 1
	}
	
	from := (request.Page-1)*request.Size + 1
	to := request.Page * request.Size
	if int64(to) > total {
		to = int(total)
	}
	if total == 0 {
		from = 0
		to = 0
	}
	
	return &model.WebResponse[[]model.OperatorResponse]{
		Data: responses,
		Meta: &model.PageMetadata{
			CurrentPage: request.Page,
			PerPage:     request.Size,
			Total:       total,
			LastPage:    lastPage,
			From:        from,
			To:          to,
		},
	}, nil
}
