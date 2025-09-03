package handler

import (
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RoleController struct {
	UseCase usecase.RoleUseCase
	Log     *logrus.Logger
}

func NewRoleController(useCase usecase.RoleUseCase, log *logrus.Logger) *RoleController {
	return &RoleController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create godoc
// @Summary Create a new role
// @Description Create a new role with name and description
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateRoleRequest true "Create role request"
// @Success 200 {object} model.SwaggerWebResponse "Role created successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles [post]
func (c *RoleController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create role")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create role", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Role created successfully", response)
}

// List godoc
// @Summary List roles
// @Description Get list of roles with optional filtering
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "Filter by role name"
// @Param is_active query bool false "Filter by active status" default(true)
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} model.SwaggerPageResponse "List of roles"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles [get]
func (c *RoleController) List(ctx *fiber.Ctx) error {
	name := ctx.Query("name", "")
	isActive := ctx.QueryBool("is_active", true)

	request := &model.ListRoleRequest{
		Name:     &name,
		IsActive: &isActive,
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list roles")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve roles", err.Error())
	}

	response := utils.SuccessResponseWithMeta("Roles retrieved successfully", responses.Data, responses.Meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Get godoc
// @Summary Get role by ID
// @Description Get role details by role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param roleId path string true "Role ID"
// @Success 200 {object} model.SwaggerWebResponse "Role details"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Role not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles/{roleId} [get]
func (c *RoleController) Get(ctx *fiber.Ctx) error {
	roleId := ctx.Params("roleId")

	request := &model.GetRoleRequest{
		ID: roleId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get role")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Role not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Role retrieved successfully", response)
}

// Update godoc
// @Summary Update role
// @Description Update role information by role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param roleId path string true "Role ID"
// @Param request body model.UpdateRoleRequest true "Update role request"
// @Success 200 {object} model.SwaggerWebResponse "Role updated successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Role not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles/{roleId} [put]
func (c *RoleController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.ID = ctx.Params("roleId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update role")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update role", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Role updated successfully", response)
}

// Delete godoc
// @Summary Delete role
// @Description Delete role by role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param roleId path string true "Role ID"
// @Success 200 {object} model.SwaggerWebResponse "Role deleted successfully"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Role not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles/{roleId} [delete]
func (c *RoleController) Delete(ctx *fiber.Ctx) error {
	roleId := ctx.Params("roleId")

	request := &model.DeleteRoleRequest{
		ID: roleId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete role")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Role not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Role deleted successfully", true)
}

// AssignPermissions godoc
// @Summary Assign permissions to role
// @Description Assign multiple permissions to a specific role
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param roleId path string true "Role ID"
// @Param request body model.AssignPermissionsRequest true "Assign permissions request"
// @Success 200 {object} model.SwaggerWebResponse "Permissions assigned successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Role not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles/{roleId}/permissions [post]
func (c *RoleController) AssignPermissions(ctx *fiber.Ctx) error {
	request := new(model.AssignPermissionsRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.RoleID = ctx.Params("roleId")

	if err := c.UseCase.AssignPermissions(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to assign permissions to role")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to assign permissions", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Permissions assigned successfully", true)
}

// RemovePermissions godoc
// @Summary Remove permissions from role
// @Description Remove multiple permissions from a specific role
// @Tags Roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param roleId path string true "Role ID"
// @Param request body model.RemovePermissionsRequest true "Remove permissions request"
// @Success 200 {object} model.SwaggerWebResponse "Permissions removed successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Role not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/roles/{roleId}/permissions [delete]
func (c *RoleController) RemovePermissions(ctx *fiber.Ctx) error {
	request := new(model.RemovePermissionsRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.RoleID = ctx.Params("roleId")

	if err := c.UseCase.RemovePermissions(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to remove permissions from role")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to remove permissions", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Permissions removed successfully", true)
}
