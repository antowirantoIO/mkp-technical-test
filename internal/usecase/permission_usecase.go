package usecase

import (
	"context"
	"mkp-boarding-test/internal/entity"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/internal/model/converter"
	"mkp-boarding-test/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PermissionUseCase struct {
	DB                       *gorm.DB
	Log                      *logrus.Logger
	Validate                 *validator.Validate
	PermissionRepository     *repository.PermissionRepository
	RolePermissionRepository *repository.RolePermissionRepository
}

func NewPermissionUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	permissionRepository *repository.PermissionRepository,
	rolePermissionRepository *repository.RolePermissionRepository) *PermissionUseCase {
	return &PermissionUseCase{
		DB:                       db,
		Log:                      logger,
		Validate:                 validate,
		PermissionRepository:     permissionRepository,
		RolePermissionRepository: rolePermissionRepository,
	}
}

func (c *PermissionUseCase) Create(ctx context.Context, request *model.CreatePermissionRequest) (*model.PermissionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	// Check if permission name already exists
	if count, err := c.PermissionRepository.CountByName(tx, request.Name, ""); err != nil {
		c.Log.WithError(err).Error("failed to count permission by name")
		return nil, fiber.ErrInternalServerError
	} else if count > 0 {
		c.Log.Error("permission name already exists")
		return nil, fiber.ErrConflict
	}

	isActive := true
	if request.IsActive != nil {
		isActive = *request.IsActive
	}

	permission := &entity.Permission{
		ID:          uuid.NewString(),
		Name:        request.Name,
		DisplayName: request.DisplayName,
		Description: request.Description,
		Resource:    request.Resource,
		Action:      request.Action,
		IsActive:    isActive,
		IsSystem:    false,
	}

	if err := c.PermissionRepository.Create(tx, permission); err != nil {
		c.Log.WithError(err).Error("failed to create permission")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PermissionToResponse(permission), nil
}

func (c *PermissionUseCase) Update(ctx context.Context, request *model.UpdatePermissionRequest) (*model.PermissionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	permission := new(entity.Permission)
	if err := c.PermissionRepository.FindById(tx, permission, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find permission")
		return nil, fiber.ErrNotFound
	}

	// Check if permission is system permission
	if permission.IsSystem {
		c.Log.Error("cannot update system permission")
		return nil, fiber.ErrForbidden
	}

	// Check if new name already exists (if name is being changed)
	if request.Name != nil && *request.Name != permission.Name {
		if count, err := c.PermissionRepository.CountByName(tx, *request.Name, permission.ID); err != nil {
			c.Log.WithError(err).Error("failed to count permission by name")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("permission name already exists")
			return nil, fiber.ErrConflict
		}
		permission.Name = *request.Name
	}

	if request.DisplayName != nil {
		permission.DisplayName = *request.DisplayName
	}
	if request.Description != nil {
		permission.Description = request.Description
	}
	if request.Resource != nil {
		permission.Resource = *request.Resource
	}
	if request.Action != nil {
		permission.Action = *request.Action
	}
	if request.IsActive != nil {
		permission.IsActive = *request.IsActive
	}

	if err := c.PermissionRepository.Update(tx, permission); err != nil {
		c.Log.WithError(err).Error("failed to update permission")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PermissionToResponse(permission), nil
}

func (c *PermissionUseCase) Get(ctx context.Context, request *model.GetPermissionRequest) (*model.PermissionResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	permission := new(entity.Permission)
	if err := c.PermissionRepository.FindById(tx, permission, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find permission")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PermissionToResponse(permission), nil
}

func (c *PermissionUseCase) Delete(ctx context.Context, request *model.DeletePermissionRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	permission := new(entity.Permission)
	if err := c.PermissionRepository.FindById(tx, permission, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find permission")
		return fiber.ErrNotFound
	}

	// Check if permission is system permission
	if permission.IsSystem {
		c.Log.Error("cannot delete system permission")
		return fiber.ErrForbidden
	}

	// Check if permission is assigned to any roles
	rolePermissions, err := c.RolePermissionRepository.FindAllByPermissionID(tx, request.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to check role permissions")
		return fiber.ErrInternalServerError
	}
	if len(rolePermissions) > 0 {
		c.Log.Error("cannot delete permission that is assigned to roles")
		return fiber.ErrConflict
	}

	if err := c.PermissionRepository.Delete(tx, permission); err != nil {
		c.Log.WithError(err).Error("failed to delete permission")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *PermissionUseCase) List(ctx context.Context, request *model.ListPermissionRequest) (*model.WebResponse[[]model.PermissionResponse], error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	query := tx.Model(&entity.Permission{}).Where("deleted_at IS NULL")

	if request.IsActive != nil {
		query = query.Where("is_active = ?", *request.IsActive)
	}
	if request.Name != nil && *request.Name != "" {
		query = query.Where("name ILIKE ?", "%"+*request.Name+"%")
	}
	if request.Resource != nil && *request.Resource != "" {
		query = query.Where("resource = ?", *request.Resource)
	}
	if request.Action != nil && *request.Action != "" {
		query = query.Where("action = ?", *request.Action)
	}

	// Count total records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.Log.WithError(err).Error("failed to count permissions")
		return nil, fiber.ErrInternalServerError
	}

	// Apply pagination
	offset := (request.Page - 1) * request.Size
	query = query.Offset(offset).Limit(request.Size)

	var permissions []entity.Permission
	if err := query.Find(&permissions).Error; err != nil {
		c.Log.WithError(err).Error("failed to find permissions")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.PermissionResponse, len(permissions))
	for i, permission := range permissions {
		responses[i] = *converter.PermissionToResponse(&permission)
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
	
	return &model.WebResponse[[]model.PermissionResponse]{
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
