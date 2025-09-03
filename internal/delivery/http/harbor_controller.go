package http

import (
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HarborController struct {
	UseCase *usecase.HarborUseCase
	Log     *logrus.Logger
}

func NewHarborController(useCase *usecase.HarborUseCase, log *logrus.Logger) *HarborController {
	return &HarborController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *HarborController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateHarborRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create harbor")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.HarborResponse]{Data: response})
}

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
		return err
	}

	return ctx.JSON(responses)
}

func (c *HarborController) Get(ctx *fiber.Ctx) error {
	harborId := ctx.Params("harborId")

	request := &model.GetHarborRequest{
		ID: harborId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get harbor")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.HarborResponse]{Data: response})
}

func (c *HarborController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateHarborRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("harborId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update harbor")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.HarborResponse]{Data: response})
}

func (c *HarborController) Delete(ctx *fiber.Ctx) error {
	harborId := ctx.Params("harborId")

	request := &model.DeleteHarborRequest{
		ID: harborId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete harbor")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}