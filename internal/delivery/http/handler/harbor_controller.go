package handler

import (
	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HarborController struct {
	UseCase usecase.HarborUseCase
	Log     *logrus.Logger
}

func NewHarborController(useCase usecase.HarborUseCase, log *logrus.Logger) *HarborController {
	return &HarborController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create godoc
// @Summary Create a new harbor
// @Description Create a new harbor with detailed information
// @Tags Harbors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.CreateHarborRequest true "Create harbor request"
// @Success 200 {object} model.SwaggerWebResponse "Harbor created successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/harbors [post]
func (c *HarborController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateHarborRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create harbor")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to create harbor", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Harbor created successfully", response)
}

// List godoc
// @Summary List harbors
// @Description Get list of harbors with optional filtering
// @Tags Harbors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param name query string false "Filter by harbor name"
// @Param country query string false "Filter by country"
// @Param province query string false "Filter by province"
// @Param city query string false "Filter by city"
// @Param is_active query bool false "Filter by active status" default(true)
// @Param page query int false "Page number" default(1)
// @Param size query int false "Page size" default(10)
// @Success 200 {object} model.SwaggerPageResponse "List of harbors"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/harbors [get]
func (c *HarborController) List(ctx *fiber.Ctx) error {
	name := ctx.Query("name", "")
	country := ctx.Query("country", "")
	province := ctx.Query("province", "")
	city := ctx.Query("city", "")
	isActive := ctx.QueryBool("is_active", true)

	request := &model.ListHarborRequest{
		Name:     &name,
		Country:  &country,
		Province: &province,
		City:     &city,
		IsActive: &isActive,
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	responses, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list harbors")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to retrieve harbors", err.Error())
	}

	response := utils.SuccessResponseWithMeta("Harbors retrieved successfully", responses.Data, responses.Meta)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// Get godoc
// @Summary Get harbor by ID
// @Description Get harbor details by harbor ID
// @Tags Harbors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param harborId path string true "Harbor ID"
// @Success 200 {object} model.SwaggerWebResponse "Harbor details"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Harbor not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/harbors/{harborId} [get]
func (c *HarborController) Get(ctx *fiber.Ctx) error {
	harborId := ctx.Params("harborId")

	request := &model.GetHarborRequest{
		ID: harborId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get harbor")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Harbor not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Harbor retrieved successfully", response)
}

// Update godoc
// @Summary Update harbor
// @Description Update harbor information by harbor ID
// @Tags Harbors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param harborId path string true "Harbor ID"
// @Param request body model.UpdateHarborRequest true "Update harbor request"
// @Success 200 {object} model.SwaggerWebResponse "Harbor updated successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Harbor not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/harbors/{harborId} [put]
func (c *HarborController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateHarborRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.ID = ctx.Params("harborId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update harbor")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update harbor", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Harbor updated successfully", response)
}

// Delete godoc
// @Summary Delete harbor
// @Description Delete harbor by harbor ID
// @Tags Harbors
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param harborId path string true "Harbor ID"
// @Success 200 {object} model.SwaggerWebResponse "Harbor deleted successfully"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 404 {object} model.SwaggerWebResponse "Harbor not found"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/harbors/{harborId} [delete]
func (c *HarborController) Delete(ctx *fiber.Ctx) error {
	harborId := ctx.Params("harborId")

	request := &model.DeleteHarborRequest{
		ID: harborId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete harbor")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "Harbor not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Harbor deleted successfully", true)
}
