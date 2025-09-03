package handler

import (
	"mkp-boarding-test/internal/delivery/http/middleware"
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OperatorController struct {
	UseCase usecase.OperatorUseCase
	Log     *logrus.Logger
}

func NewOperatorController(useCase usecase.OperatorUseCase, log *logrus.Logger) *OperatorController {
	return &OperatorController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create godoc
// @Summary Create a new operator
// @Description Create a new operator with company information
// @Tags Operators
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateOperatorRequest true "Create operator request"
// @Success 200 {object} model.SwaggerWebResponse "Operator created successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/operators [post]
func (c *OperatorController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateOperatorRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.UserID = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create operator")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create operator", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Operator created successfully", response)
}

// List godoc
// @Summary List operators
// @Description Get list of operators with optional filtering
// @Tags Operators
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param company_name query string false "Filter by company name"
// @Param operator_type query string false "Filter by operator type"
// @Param status query string false "Filter by status"
// @Param country query string false "Filter by country"
// @Param province query string false "Filter by province"
// @Param city query string false "Filter by city"
// @Param is_active query bool false "Filter by active status" default(true)
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} model.SwaggerPageResponse "List of operators"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/operators [get]
func (c *OperatorController) List(ctx *fiber.Ctx) error {
	companyName := ctx.Query("company_name", "")
	operatorType := ctx.Query("operator_type", "")
	status := ctx.Query("status", "")
	country := ctx.Query("country", "")
	province := ctx.Query("province", "")
	city := ctx.Query("city", "")
	isActive := ctx.QueryBool("is_active", true)

	request := &model.ListOperatorRequest{
		CompanyName:  &companyName,
		OperatorType: &operatorType,
		Status:       &status,
		Country:      &country,
		Province:     &province,
		City:         &city,
		IsActive:     &isActive,
		Page:         ctx.QueryInt("page", 1),
		Size:         ctx.QueryInt("size", 10),
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list operators")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve operators", err.Error())
	}

	response := utils.SuccessResponseWithMeta("Operators retrieved successfully", responses.Data, responses.Meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Get godoc
// @Summary Get operator by ID
// @Description Get operator details by operator ID
// @Tags Operators
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param operatorId path string true "Operator ID"
// @Success 200 {object} model.SwaggerWebResponse "Operator details"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Operator not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/operators/{operatorId} [get]
func (c *OperatorController) Get(ctx *fiber.Ctx) error {
	operatorId := ctx.Params("operatorId")

	request := &model.GetOperatorRequest{
		ID: operatorId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get operator")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Operator not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Operator retrieved successfully", response)
}

// Update godoc
// @Summary Update operator
// @Description Update operator information by operator ID
// @Tags Operators
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param operatorId path string true "Operator ID"
// @Param request body model.UpdateOperatorRequest true "Update operator request"
// @Success 200 {object} model.SwaggerWebResponse "Operator updated successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Operator not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/operators/{operatorId} [put]
func (c *OperatorController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateOperatorRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.ID = ctx.Params("operatorId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update operator")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update operator", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Operator updated successfully", response)
}

// Delete godoc
// @Summary Delete operator
// @Description Delete operator by operator ID
// @Tags Operators
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param operatorId path string true "Operator ID"
// @Success 200 {object} model.SwaggerWebResponse "Operator deleted successfully"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Operator not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/operators/{operatorId} [delete]
func (c *OperatorController) Delete(ctx *fiber.Ctx) error {
	operatorId := ctx.Params("operatorId")

	request := &model.DeleteOperatorRequest{
		ID: operatorId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete operator")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Operator not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Operator deleted successfully", true)
}
