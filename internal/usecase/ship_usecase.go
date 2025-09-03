package usecase

import (
	"context"
	"mkp-boarding-test/internal/converter"
	"mkp-boarding-test/internal/entity"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShipUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	ShipRepository *repository.ShipRepository
}

func NewShipUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, shipRepository *repository.ShipRepository) *ShipUseCase {
	return &ShipUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		ShipRepository: shipRepository,
	}
}

func (c *ShipUseCase) Create(ctx context.Context, request *model.CreateShipRequest) (*model.ShipResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	// Check if ship name already exists for the operator
	if count, err := c.ShipRepository.CountByShipNameAndOperatorID(tx, request.ShipName, request.OperatorID, ""); err != nil {
		c.Log.WithError(err).Error("failed to count ship by name and operator")
		return nil, fiber.ErrInternalServerError
	} else if count > 0 {
		c.Log.Error("ship name already exists for this operator")
		return nil, fiber.ErrConflict
	}

	// Check if IMO number already exists (if provided)
	if request.IMONumber != nil && *request.IMONumber != "" {
		if count, err := c.ShipRepository.CountByIMONumber(tx, *request.IMONumber, ""); err != nil {
			c.Log.WithError(err).Error("failed to count ship by IMO number")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("IMO number already exists")
			return nil, fiber.ErrConflict
		}
	}

	imoNumber := ""
	if request.IMONumber != nil {
		imoNumber = *request.IMONumber
	}
	callSign := ""
	if request.CallSign != nil {
		callSign = *request.CallSign
	}
	mmsi := ""
	if request.MMSI != nil {
		mmsi = *request.MMSI
	}

	ship := &entity.Ship{
		ID:                    uuid.NewString(),
		OperatorID:            request.OperatorID,
		ShipName:              request.ShipName,
		IMONumber:             imoNumber,
		CallSign:              callSign,
		MMSI:                  mmsi,
		ShipType:              request.ShipType,
		FlagState:             request.FlagState,
		PortOfRegistry:        request.PortOfRegistry,
		BuildYear:             request.BuildYear,
		Builder:               request.Builder,
		Length:                request.Length,
		Beam:                  request.Beam,
		Draft:                 request.Draft,
		GrossTonnage:          request.GrossTonnage,
		NetTonnage:            request.NetTonnage,
		DeadweightTonnage:     request.DeadweightTonnage,
		MaxSpeed:              request.MaxSpeed,
		PassengerCapacity:     request.PassengerCapacity,
		CrewCapacity:          request.CrewCapacity,
		ClassificationSociety: request.ClassificationSociety,
		LastInspection:        request.LastInspection,
		NextInspection:        request.NextInspection,
		InsuranceExpiry:       request.InsuranceExpiry,
		CertificateExpiry:     request.CertificateExpiry,
		CurrentLatitude:       request.CurrentLatitude,
		CurrentLongitude:      request.CurrentLongitude,
		Status:                "active",
		IsActive:              true,
		Notes:                 request.Notes,
	}

	if err := c.ShipRepository.Create(tx, ship); err != nil {
		c.Log.WithError(err).Error("failed to create ship")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ShipToResponse(ship), nil
}

func (c *ShipUseCase) Update(ctx context.Context, request *model.UpdateShipRequest) (*model.ShipResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	ship := &entity.Ship{}
	if err := c.ShipRepository.FindById(tx, ship, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find ship")
		return nil, fiber.ErrNotFound
	}

	// Check if ship name already exists for the operator (exclude current ship)
	if request.ShipName != nil {
		if count, err := c.ShipRepository.CountByShipNameAndOperatorID(tx, *request.ShipName, ship.OperatorID, ship.ID); err != nil {
			c.Log.WithError(err).Error("failed to count ship by name and operator")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("ship name already exists for this operator")
			return nil, fiber.ErrConflict
		}
		ship.ShipName = *request.ShipName
	}

	// Check if IMO number already exists (exclude current ship)
	if request.IMONumber != nil && *request.IMONumber != "" {
		if count, err := c.ShipRepository.CountByIMONumber(tx, *request.IMONumber, ship.ID); err != nil {
			c.Log.WithError(err).Error("failed to count ship by IMO number")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("IMO number already exists")
			return nil, fiber.ErrConflict
		}
		ship.IMONumber = *request.IMONumber
	}

	if request.CallSign != nil {
		ship.CallSign = *request.CallSign
	}
	if request.MMSI != nil {
		ship.MMSI = *request.MMSI
	}
	if request.ShipType != nil {
		ship.ShipType = *request.ShipType
	}
	if request.FlagState != nil {
		ship.FlagState = *request.FlagState
	}
	if request.PortOfRegistry != nil {
		ship.PortOfRegistry = *request.PortOfRegistry
	}
	if request.BuildYear != nil {
		ship.BuildYear = request.BuildYear
	}
	if request.Builder != nil {
		ship.Builder = request.Builder
	}
	if request.Length != nil {
		ship.Length = request.Length
	}
	if request.Beam != nil {
		ship.Beam = request.Beam
	}
	if request.Draft != nil {
		ship.Draft = request.Draft
	}
	if request.GrossTonnage != nil {
		ship.GrossTonnage = request.GrossTonnage
	}
	if request.NetTonnage != nil {
		ship.NetTonnage = request.NetTonnage
	}
	if request.DeadweightTonnage != nil {
		ship.DeadweightTonnage = request.DeadweightTonnage
	}
	if request.MaxSpeed != nil {
		ship.MaxSpeed = request.MaxSpeed
	}
	if request.PassengerCapacity != nil {
		ship.PassengerCapacity = request.PassengerCapacity
	}
	if request.CrewCapacity != nil {
		ship.CrewCapacity = request.CrewCapacity
	}
	if request.ClassificationSociety != nil {
		ship.ClassificationSociety = request.ClassificationSociety
	}
	if request.LastInspection != nil {
		ship.LastInspection = request.LastInspection
	}
	if request.NextInspection != nil {
		ship.NextInspection = request.NextInspection
	}
	if request.InsuranceExpiry != nil {
		ship.InsuranceExpiry = request.InsuranceExpiry
	}
	if request.CertificateExpiry != nil {
		ship.CertificateExpiry = request.CertificateExpiry
	}
	if request.CurrentLatitude != nil {
		ship.CurrentLatitude = request.CurrentLatitude
	}
	if request.CurrentLongitude != nil {
		ship.CurrentLongitude = request.CurrentLongitude
	}
	if request.Status != nil {
		ship.Status = *request.Status
	}
	if request.IsActive != nil {
		ship.IsActive = *request.IsActive
	}
	if request.Notes != nil {
		ship.Notes = request.Notes
	}

	if err := c.ShipRepository.Update(tx, ship); err != nil {
		c.Log.WithError(err).Error("failed to update ship")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ShipToResponse(ship), nil
}

func (c *ShipUseCase) Get(ctx context.Context, request *model.GetShipRequest) (*model.ShipResponse, error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	ship := &entity.Ship{}
	if err := c.ShipRepository.FindById(tx, ship, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find ship")
		return nil, fiber.ErrNotFound
	}

	return converter.ShipToResponse(ship), nil
}

func (c *ShipUseCase) Delete(ctx context.Context, request *model.DeleteShipRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	ship := &entity.Ship{}
	if err := c.ShipRepository.FindById(tx, ship, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find ship")
		return fiber.ErrNotFound
	}

	if err := c.ShipRepository.Delete(tx, ship); err != nil {
		c.Log.WithError(err).Error("failed to delete ship")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *ShipUseCase) List(ctx context.Context, request *model.ListShipRequest) (*model.WebResponse[[]model.ShipResponse], error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	query := tx.Model(&entity.Ship{}).Where("deleted_at IS NULL")

	if request.OperatorID != nil && *request.OperatorID != "" {
		query = query.Where("operator_id = ?", *request.OperatorID)
	}
	if request.IsActive != nil {
		query = query.Where("is_active = ?", *request.IsActive)
	}
	if request.ShipName != nil && *request.ShipName != "" {
		query = query.Where("ship_name ILIKE ?", "%"+*request.ShipName+"%")
	}
	if request.ShipType != nil && *request.ShipType != "" {
		query = query.Where("ship_type = ?", *request.ShipType)
	}
	if request.FlagState != nil && *request.FlagState != "" {
		query = query.Where("flag_state = ?", *request.FlagState)
	}
	if request.Status != nil && *request.Status != "" {
		query = query.Where("status = ?", *request.Status)
	}

	// Count total records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.Log.WithError(err).Error("failed to count ships")
		return nil, fiber.ErrInternalServerError
	}

	// Apply pagination
	offset := (request.Page - 1) * request.Size
	query = query.Offset(offset).Limit(request.Size)

	var ships []entity.Ship
	if err := query.Find(&ships).Error; err != nil {
		c.Log.WithError(err).Error("failed to find ships")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.ShipResponse, len(ships))
	for i, ship := range ships {
		responses[i] = *converter.ShipToResponse(&ship)
	}

	return &model.WebResponse[[]model.ShipResponse]{
		Data: responses,
		Paging: &model.PageMetadata{
			Page:      request.Page,
			Size:      request.Size,
			TotalItem: total,
			TotalPage: (total + int64(request.Size) - 1) / int64(request.Size),
		},
	}, nil
}
