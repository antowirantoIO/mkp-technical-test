package http

import (
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PermissionController struct {
	UseCase *usecase.PermissionUseCase
	Log     *logrus.Logger
}

func NewPermissionController(useCase *usecase.PermissionUseCase, log *logrus.Logger) *PermissionController {
	return &PermissionController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *PermissionController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreatePermissionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create permission")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.PermissionResponse]{Data: response})
}

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
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.PermissionResponse]{Data: responses})
}

func (c *PermissionController) Get(ctx *fiber.Ctx) error {
	permissionId := ctx.Params("permissionId")

	request := &model.GetPermissionRequest{
		ID: permissionId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get permission")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.PermissionResponse]{Data: response})
}

func (c *PermissionController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdatePermissionRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("permissionId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update permission")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.PermissionResponse]{Data: response})
}

func (c *PermissionController) Delete(ctx *fiber.Ctx) error {
	permissionId := ctx.Params("permissionId")

	request := &model.DeletePermissionRequest{
		ID: permissionId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete permission")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}