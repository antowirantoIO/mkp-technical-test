package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type OperatorController struct {
	UseCase *usecase.OperatorUseCase
	Log     *logrus.Logger
}

func NewOperatorController(useCase *usecase.OperatorUseCase, log *logrus.Logger) *OperatorController {
	return &OperatorController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *OperatorController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateOperatorRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.UserID = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create operator")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.OperatorResponse]{Data: response})
}

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
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.OperatorResponse]{Data: responses})
}

func (c *OperatorController) Get(ctx *fiber.Ctx) error {
	operatorId := ctx.Params("operatorId")

	request := &model.GetOperatorRequest{
		ID: operatorId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get operator")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.OperatorResponse]{Data: response})
}

func (c *OperatorController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateOperatorRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("operatorId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update operator")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.OperatorResponse]{Data: response})
}

func (c *OperatorController) Delete(ctx *fiber.Ctx) error {
	operatorId := ctx.Params("operatorId")

	request := &model.DeleteOperatorRequest{
		ID: operatorId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete operator")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}