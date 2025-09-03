package http

import (
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ShipController struct {
	UseCase *usecase.ShipUseCase
	Log     *logrus.Logger
}

func NewShipController(useCase *usecase.ShipUseCase, log *logrus.Logger) *ShipController {
	return &ShipController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *ShipController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateShipRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create ship")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ShipResponse]{Data: response})
}

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
		return err
	}

	return ctx.JSON(responses)
}

func (c *ShipController) Get(ctx *fiber.Ctx) error {
	shipId := ctx.Params("shipId")

	request := &model.GetShipRequest{
		ID: shipId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get ship")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ShipResponse]{Data: response})
}

func (c *ShipController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateShipRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("shipId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update ship")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ShipResponse]{Data: response})
}

func (c *ShipController) Delete(ctx *fiber.Ctx) error {
	shipId := ctx.Params("shipId")

	request := &model.DeleteShipRequest{
		ID: shipId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete ship")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}