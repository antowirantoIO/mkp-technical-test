package handler

import (
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ShipController struct {
	UseCase usecase.ShipUseCase
	Log     *logrus.Logger
}

func NewShipController(useCase usecase.ShipUseCase, log *logrus.Logger) *ShipController {
	return &ShipController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create godoc
// @Summary Create a new ship
// @Description Create a new ship with detailed information
// @Tags Ships
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateShipRequest true "Create ship request"
// @Success 200 {object} model.SwaggerWebResponse "Ship created successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/ships [post]
func (c *ShipController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateShipRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create ship")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create ship", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Ship created successfully", response)
}

// List godoc
// @Summary List ships
// @Description Get list of ships with optional filtering
// @Tags Ships
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param operator_id query string false "Filter by operator ID"
// @Param ship_name query string false "Filter by ship name"
// @Param flag_state query string false "Filter by flag state"
// @Param ship_type query string false "Filter by ship type"
// @Param status query string false "Filter by status"
// @Param is_active query bool false "Filter by active status" default(true)
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} model.SwaggerPageResponse "List of ships"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/ships [get]
func (c *ShipController) List(ctx *fiber.Ctx) error {
	operatorID := ctx.Query("operator_id", "")
	shipName := ctx.Query("ship_name", "")
	flagState := ctx.Query("flag_state", "")
	shipType := ctx.Query("ship_type", "")
	status := ctx.Query("status", "")
	isActive := ctx.QueryBool("is_active", true)

	request := &model.ListShipRequest{
		OperatorID: &operatorID,
		ShipName:   &shipName,
		FlagState:  &flagState,
		ShipType:   &shipType,
		Status:     &status,
		IsActive:   &isActive,
		Page:       ctx.QueryInt("page", 1),
		Size:       ctx.QueryInt("size", 10),
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list ships")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve ships", err.Error())
	}

	response := utils.SuccessResponseWithMeta("Ships retrieved successfully", responses.Data, responses.Meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Get godoc
// @Summary Get ship by ID
// @Description Get ship details by ship ID
// @Tags Ships
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param shipId path string true "Ship ID"
// @Success 200 {object} model.SwaggerWebResponse "Ship details"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Ship not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/ships/{shipId} [get]
func (c *ShipController) Get(ctx *fiber.Ctx) error {
	shipId := ctx.Params("shipId")

	request := &model.GetShipRequest{
		ID: shipId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get ship")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Ship not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Ship retrieved successfully", response)
}

// Update godoc
// @Summary Update ship
// @Description Update ship information by ship ID
// @Tags Ships
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param shipId path string true "Ship ID"
// @Param request body model.UpdateShipRequest true "Update ship request"
// @Success 200 {object} model.SwaggerWebResponse "Ship updated successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Ship not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/ships/{shipId} [put]
func (c *ShipController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateShipRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.ID = ctx.Params("shipId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update ship")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update ship", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Ship updated successfully", response)
}

// Delete godoc
// @Summary Delete ship
// @Description Delete ship by ship ID
// @Tags Ships
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param shipId path string true "Ship ID"
// @Success 200 {object} model.SwaggerWebResponse "Ship deleted successfully"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Ship not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/ships/{shipId} [delete]
func (c *ShipController) Delete(ctx *fiber.Ctx) error {
	shipId := ctx.Params("shipId")

	request := &model.DeleteShipRequest{
		ID: shipId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete ship")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Ship not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Ship deleted successfully", true)
}
