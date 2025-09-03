package role

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

type RoleUseCaseImpl struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Validate             *validator.Validate
	RoleRepository       repository.RoleRepository
	PermissionRepository repository.PermissionRepository
}

func NewRoleUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	roleRepository repository.RoleRepository, permissionRepository repository.PermissionRepository) usecase.RoleUseCase {
	return &RoleUseCaseImpl{
		DB:                   db,
		Log:                  logger,
		Validate:             validate,
		RoleRepository:       roleRepository,
		PermissionRepository: permissionRepository,
	}
}

func (c *RoleUseCaseImpl) Create(ctx context.Context, request *model.CreateRoleRequest) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	// Check if role name already exists
	if count, err := c.RoleRepository.CountByName(tx, request.Name, ""); err != nil {
		c.Log.WithError(err).Error("failed to count role by name")
		return nil, fiber.ErrInternalServerError
	} else if count > 0 {
		c.Log.Error("role name already exists")
		return nil, fiber.ErrConflict
	}

	isActive := true
	if request.IsActive != nil {
		isActive = *request.IsActive
	}

	role := &entity.Role{
		ID:          uuid.NewString(),
		Name:        request.Name,
		DisplayName: request.DisplayName,
		Description: request.Description,
		IsActive:    isActive,
		IsSystem:    false,
	}

	if err := c.RoleRepository.Create(tx, role); err != nil {
		c.Log.WithError(err).Error("failed to create role")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleToResponse(role), nil
}

func (c *RoleUseCaseImpl) Update(ctx context.Context, request *model.UpdateRoleRequest) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find role")
		return nil, fiber.ErrNotFound
	}

	// Check if role is system role
	if role.IsSystem {
		c.Log.Error("cannot update system role")
		return nil, fiber.ErrForbidden
	}

	// Check if new name already exists (if name is being changed)
	if request.Name != nil && *request.Name != role.Name {
		if count, err := c.RoleRepository.CountByName(tx, *request.Name, role.ID); err != nil {
			c.Log.WithError(err).Error("failed to count role by name")
			return nil, fiber.ErrInternalServerError
		} else if count > 0 {
			c.Log.Error("role name already exists")
			return nil, fiber.ErrConflict
		}
		role.Name = *request.Name
	}

	if request.DisplayName != nil {
		role.DisplayName = *request.DisplayName
	}
	if request.Description != nil {
		role.Description = request.Description
	}
	if request.IsActive != nil {
		role.IsActive = *request.IsActive
	}

	if err := c.RoleRepository.Update(tx, role); err != nil {
		c.Log.WithError(err).Error("failed to update role")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleToResponse(role), nil
}

func (c *RoleUseCaseImpl) Get(ctx context.Context, request *model.GetRoleRequest) (*model.RoleResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find role")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoleToResponse(role), nil
}

func (c *RoleUseCaseImpl) Delete(ctx context.Context, request *model.DeleteRoleRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.ID); err != nil {
		c.Log.WithError(err).Error("failed to find role")
		return fiber.ErrNotFound
	}

	// Check if role is system role
	if role.IsSystem {
		c.Log.Error("cannot delete system role")
		return fiber.ErrForbidden
	}

	// Delete all role permissions first
	// TODO: Implement role permission deletion when RolePermissionRepository is available
	// if err := c.RolePermissionRepository.DeleteByRoleID(tx, role.ID); err != nil {
	//	c.Log.WithError(err).Error("failed to delete role permissions")
	//	return fiber.ErrInternalServerError
	// }

	if err := c.RoleRepository.Delete(tx, role); err != nil {
		c.Log.WithError(err).Error("failed to delete role")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *RoleUseCaseImpl) List(ctx context.Context, request *model.ListRoleRequest) (*model.WebResponse[[]model.RoleResponse], error) {
	tx := c.DB.WithContext(ctx)

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	query := tx.Model(&entity.Role{}).Where("deleted_at IS NULL")

	if request.IsActive != nil {
		query = query.Where("is_active = ?", *request.IsActive)
	}
	if request.Name != nil && *request.Name != "" {
		query = query.Where("name ILIKE ?", "%"+*request.Name+"%")
	}

	// Count total records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.Log.WithError(err).Error("failed to count roles")
		return nil, fiber.ErrInternalServerError
	}

	// Apply pagination
	offset := (request.Page - 1) * request.Size
	query = query.Offset(offset).Limit(request.Size)

	var roles []entity.Role
	if err := query.Find(&roles).Error; err != nil {
		c.Log.WithError(err).Error("failed to find roles")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		responses[i] = *converter.RoleToResponse(&role)
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

	return &model.WebResponse[[]model.RoleResponse]{
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

func (c *RoleUseCaseImpl) AssignPermissions(ctx context.Context, request *model.AssignPermissionsRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	// Check if role exists
	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.RoleID); err != nil {
		c.Log.WithError(err).Error("failed to find role")
		return fiber.ErrNotFound
	}

	// Check if role is system role
	if role.IsSystem {
		c.Log.Error("cannot modify permissions for system role")
		return fiber.ErrForbidden
	}

	// Validate all permissions exist
	for _, permissionID := range request.PermissionIDs {
		permission := new(entity.Permission)
		if err := c.PermissionRepository.FindById(tx, permission, permissionID); err != nil {
			c.Log.WithError(err).Error("failed to find permission")
			return fiber.ErrNotFound
		}
	}

	// TODO: Create role permissions when RolePermissionRepository is available
	// for _, permissionID := range request.PermissionIDs {
	//	// Check if already exists and create role permission
	// }

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *RoleUseCaseImpl) RemovePermissions(ctx context.Context, request *model.RemovePermissionsRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return fiber.ErrBadRequest
	}

	// Check if role exists
	role := new(entity.Role)
	if err := c.RoleRepository.FindById(tx, role, request.RoleID); err != nil {
		c.Log.WithError(err).Error("failed to find role")
		return fiber.ErrNotFound
	}

	// Check if role is system role
	if role.IsSystem {
		c.Log.Error("cannot modify permissions for system role")
		return fiber.ErrForbidden
	}

	// TODO: Remove role permissions when RolePermissionRepository is available
	// for _, permissionID := range request.PermissionIDs {
	//	// Delete role permission by role ID and permission ID
	// }

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}
