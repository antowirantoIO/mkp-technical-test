package harbor

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

type HarborUseCaseImpl struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	HarborRepository repository.HarborRepository
}

func NewHarborUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, harborRepository repository.HarborRepository) usecase.HarborUseCase {
	return &HarborUseCaseImpl{
		DB:               db,
		Log:              log,
		Validate:         validate,
		HarborRepository: harborRepository,
	}
}

func (c *HarborUseCaseImpl) Create(ctx context.Context, request *model.CreateHarborRequest) (*model.HarborResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	// Check if harbor code already exists
	if count, err := c.HarborRepository.CountByHarborCode(tx, request.HarborCode, ""); err != nil {
		c.Log.WithError(err).Error("failed to count harbor by code")
		return nil, fiber.ErrInternalServerError
	} else if count > 0 {
		c.Log.Error("harbor code already exists")
		return nil, fiber.ErrConflict
	}

	// Check if UN/LOCODE already exists (if provided)
	if request.UNLocode != nil && *request.UNLocode != "" {
		if count, err := c.HarborRepository.CountByUNLocode(tx, *request.UNLocode, ""); err != nil {
			c.Log.WithError(err).Error("failed to count harbor by UN/LOCODE")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("UN/LOCODE already exists")
			return nil, fiber.ErrConflict
		}
	}

	harbor := &entity.Harbor{
		ID:         uuid.NewString(),
		HarborCode: request.HarborCode,
		HarborName: request.HarborName,
		Country:    request.Country,
		Province:   request.Province,
		City:       request.City,
		Status:     "active",
		IsActive:   true,
	}

	// Handle optional fields
	if request.UNLocode != nil {
		harbor.UNLocode = *request.UNLocode
	}
	if request.Latitude != nil {
		harbor.Latitude = *request.Latitude
	}
	if request.Longitude != nil {
		harbor.Longitude = *request.Longitude
	}
	if request.TimeZone != nil {
		harbor.Timezone = *request.TimeZone
	}
	if request.MaxShipLength != nil {
		harbor.MaxShipLength = request.MaxShipLength
	}
	if request.MaxShipBeam != nil {
		harbor.MaxShipBeam = request.MaxShipBeam
	}
	if request.MaxShipDraft != nil {
		harbor.MaxShipDraft = request.MaxShipDraft
	}
	if request.BerthCount != nil {
		harbor.BerthCount = *request.BerthCount
	}
	if request.StorageCapacity != nil {
		harbor.StorageCapacity = request.StorageCapacity
	}
	if request.ContactPerson != nil {
		harbor.ContactPerson = *request.ContactPerson
	}
	if request.ContactPhone != nil {
		harbor.ContactPhone = *request.ContactPhone
	}
	if request.ContactEmail != nil {
		harbor.ContactEmail = *request.ContactEmail
	}
	if request.Website != nil {
		harbor.Website = request.Website
	}
	if request.OperatingHours != nil {
		harbor.WorkingHours = *request.OperatingHours
	}
	if request.HasPilotage != nil {
		harbor.HasPilotage = *request.HasPilotage
	}
	if request.HasTugService != nil {
		harbor.HasTugService = *request.HasTugService
	}
	if request.HasQuarantine != nil {
		harbor.HasQuarantine = *request.HasQuarantine
	}
	if request.HasCustoms != nil {
		harbor.HasCustoms = *request.HasCustoms
	}
	if request.HasRepair != nil {
		harbor.HasRepairService = *request.HasRepair
	}
	if request.HasWaste != nil {
		harbor.HasWaste = *request.HasWaste
	}
	if request.Notes != nil {
		harbor.Notes = request.Notes
	}

	if err := c.HarborRepository.Create(tx, harbor); err != nil {
		c.Log.WithError(err).Error("failed to create harbor")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.HarborToResponse(harbor), nil
}

func (c *HarborUseCaseImpl) Update(ctx context.Context, request *model.UpdateHarborRequest) (*model.HarborResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	harbor := &entity.Harbor{}
	if err := c.HarborRepository.FindById(tx, harbor, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find harbor")
		return nil, fiber.ErrNotFound
	}

	// Check if harbor code already exists (exclude current harbor)
	if request.HarborCode != nil {
		if count, err := c.HarborRepository.CountByHarborCode(tx, *request.HarborCode, harbor.ID); err != nil {
			c.Log.WithError(err).Error("failed to count harbor by code")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("harbor code already exists")
			return nil, fiber.ErrConflict
		}
		harbor.HarborCode = *request.HarborCode
	}

	// Check if UN/LOCODE already exists (exclude current harbor)
	if request.UNLocode != nil && *request.UNLocode != "" {
		if count, err := c.HarborRepository.CountByUNLocode(tx, *request.UNLocode, harbor.ID); err != nil {
			c.Log.WithError(err).Error("failed to count harbor by UN/LOCODE")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("UN/LOCODE already exists")
			return nil, fiber.ErrConflict
		}
		harbor.UNLocode = *request.UNLocode
	}

	if request.HarborName != nil {
		harbor.HarborName = *request.HarborName
	}
	if request.Country != nil {
		harbor.Country = *request.Country
	}
	if request.City != nil {
		harbor.City = *request.City
	}
	if request.Province != nil {
		harbor.Province = *request.Province
	}
	if request.Latitude != nil {
		harbor.Latitude = *request.Latitude
	}
	if request.Longitude != nil {
		harbor.Longitude = *request.Longitude
	}
	if request.TimeZone != nil {
		harbor.Timezone = *request.TimeZone
	}
	if request.MaxShipLength != nil {
		harbor.MaxShipLength = request.MaxShipLength
	}
	if request.MaxShipBeam != nil {
		harbor.MaxShipBeam = request.MaxShipBeam
	}
	if request.MaxShipDraft != nil {
		harbor.MaxShipDraft = request.MaxShipDraft
	}
	if request.BerthCount != nil {
		harbor.BerthCount = *request.BerthCount
	}
	if request.StorageCapacity != nil {
		harbor.StorageCapacity = request.StorageCapacity
	}
	if request.HasTugService != nil {
		harbor.HasTugService = *request.HasTugService
	}
	if request.HasPilotage != nil {
		harbor.HasPilotage = *request.HasPilotage
	}
	if request.HasQuarantine != nil {
		harbor.HasQuarantine = *request.HasQuarantine
	}
	if request.HasCustoms != nil {
		harbor.HasCustoms = *request.HasCustoms
	}
	if request.HasRepair != nil {
		harbor.HasRepairService = *request.HasRepair
	}
	if request.HasWaste != nil {
		harbor.HasWaste = *request.HasWaste
	}
	if request.ContactPerson != nil {
		harbor.ContactPerson = *request.ContactPerson
	}
	if request.ContactPhone != nil {
		harbor.ContactPhone = *request.ContactPhone
	}
	if request.ContactEmail != nil {
		harbor.ContactEmail = *request.ContactEmail
	}
	if request.Website != nil {
		harbor.Website = request.Website
	}
	if request.OperatingHours != nil {
		harbor.WorkingHours = *request.OperatingHours
	}
	if request.IsActive != nil {
		harbor.IsActive = *request.IsActive
	}
	if request.Notes != nil {
		harbor.Notes = request.Notes
	}

	if err := c.HarborRepository.Update(tx, harbor); err != nil {
		c.Log.WithError(err).Error("failed to update harbor")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.HarborToResponse(harbor), nil
}

func (c *HarborUseCaseImpl) Get(ctx context.Context, request *model.GetHarborRequest) (*model.HarborResponse, error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	harbor := &entity.Harbor{}
	if err := c.HarborRepository.FindById(tx, harbor, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find harbor")
		return nil, fiber.ErrNotFound
	}

	return converter.HarborToResponse(harbor), nil
}

func (c *HarborUseCaseImpl) Delete(ctx context.Context, request *model.DeleteHarborRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	harbor := &entity.Harbor{}
	if err := c.HarborRepository.FindById(tx, harbor, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find harbor")
		return fiber.ErrNotFound
	}

	if err := c.HarborRepository.Delete(tx, harbor); err != nil {
		c.Log.WithError(err).Error("failed to delete harbor")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *HarborUseCaseImpl) List(ctx context.Context, request *model.ListHarborRequest, userId string) (*model.WebResponse[[]model.HarborResponse], error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	var userRole entity.UserRole
	err := tx.Model(&entity.UserRole{}).Where("user_id = ?", userId).First(&userRole).Error

	if err != nil {
		c.Log.WithError(err).Error("failed to find user role")
		return nil, fiber.ErrInternalServerError
	}

	query := tx.Model(&entity.Harbor{}).Where("deleted_at IS NULL")

	if request.IsActive != nil {
		query = query.Where("is_active = ?", *request.IsActive)
	}
	if request.Name != nil && *request.Name != "" {
		query = query.Where("harbor_name ILIKE ?", "%"+*request.Name+"%")
	}
	if request.Country != nil && *request.Country != "" {
		query = query.Where("country = ?", *request.Country)
	}
	if request.Province != nil && *request.Province != "" {
		query = query.Where("province = ?", *request.Province)
	}
	if request.City != nil && *request.City != "" {
		query = query.Where("city = ?", *request.City)
	}

	// Count total records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.Log.WithError(err).Error("failed to count harbors")
		return nil, fiber.ErrInternalServerError
	}

	query = query.Joins("INNER JOIN role_harbors AS rh ON rh.harbor_id = harbors.id").Where("rh.role_id = ?", userRole.RoleID)

	// Apply pagination
	offset := (request.Page - 1) * request.Size
	query = query.Offset(offset).Limit(request.Size)

	var harbors []entity.Harbor
	if err := query.Find(&harbors).Error; err != nil {
		c.Log.WithError(err).Error("failed to find harbors")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.HarborResponse, len(harbors))
	for i, harbor := range harbors {
		responses[i] = *converter.HarborToResponse(&harbor)
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

	return &model.WebResponse[[]model.HarborResponse]{
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
