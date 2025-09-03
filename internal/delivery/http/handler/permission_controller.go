package handler

import (
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PermissionController struct {
	UseCase usecase.PermissionUseCase
	Log     *logrus.Logger
}

func NewPermissionController(useCase usecase.PermissionUseCase, log *logrus.Logger) *PermissionController {
	return &PermissionController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create godoc
// @Summary Create a new permission
// @Description Create a new permission with name, resource, and action
// @Tags Permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreatePermissionRequest true "Create permission request"
// @Success 200 {object} model.SwaggerWebResponse "Permission created successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/permissions [post]
func (c *PermissionController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreatePermissionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create permission")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create permission", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Permission created successfully", response)
}

// List godoc
// @Summary List permissions
// @Description Get list of permissions with optional filtering
// @Tags Permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "Filter by permission name"
// @Param resource query string false "Filter by resource"
// @Param action query string false "Filter by action"
// @Param is_active query bool false "Filter by active status" default(true)
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} model.SwaggerPageResponse "List of permissions"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/permissions [get]
func (c *PermissionController) List(ctx *fiber.Ctx) error {
	name := ctx.Query("name", "")
	resource := ctx.Query("resource", "")
	action := ctx.Query("action", "")
	isActive := ctx.QueryBool("is_active", true)

	request := &model.ListPermissionRequest{
		Name:     &name,
		Resource: &resource,
		Action:   &action,
		IsActive: &isActive,
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list permissions")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve permissions", err.Error())
	}

	response := utils.SuccessResponseWithMeta("Permissions retrieved successfully", responses.Data, responses.Meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Get godoc
// @Summary Get permission by ID
// @Description Get permission details by permission ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param permissionId path string true "Permission ID"
// @Success 200 {object} model.SwaggerWebResponse "Permission details"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Permission not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/permissions/{permissionId} [get]
func (c *PermissionController) Get(ctx *fiber.Ctx) error {
	permissionId := ctx.Params("permissionId")

	request := &model.GetPermissionRequest{
		ID: permissionId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get permission")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Permission not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Permission retrieved successfully", response)
}

// Update godoc
// @Summary Update permission
// @Description Update permission information by permission ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param permissionId path string true "Permission ID"
// @Param request body model.UpdatePermissionRequest true "Update permission request"
// @Success 200 {object} model.SwaggerWebResponse "Permission updated successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Permission not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/permissions/{permissionId} [put]
func (c *PermissionController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdatePermissionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.ID = ctx.Params("permissionId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update permission")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update permission", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Permission updated successfully", response)
}

// Delete godoc
// @Summary Delete permission
// @Description Delete permission by permission ID
// @Tags Permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param permissionId path string true "Permission ID"
// @Success 200 {object} model.SwaggerWebResponse "Permission deleted successfully"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Permission not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/permissions/{permissionId} [delete]
func (c *PermissionController) Delete(ctx *fiber.Ctx) error {
	permissionId := ctx.Params("permissionId")

	request := &model.DeletePermissionRequest{
		ID: permissionId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete permission")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Permission not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Permission deleted successfully", true)
}
