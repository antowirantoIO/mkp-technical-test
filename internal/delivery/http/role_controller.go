package http

import (
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RoleController struct {
	UseCase *usecase.RoleUseCase
	Log     *logrus.Logger
}

func NewRoleController(useCase *usecase.RoleUseCase, log *logrus.Logger) *RoleController {
	return &RoleController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *RoleController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create role")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

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
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.RoleResponse]{Data: responses})
}

func (c *RoleController) Get(ctx *fiber.Ctx) error {
	roleId := ctx.Params("roleId")

	request := &model.GetRoleRequest{
		ID: roleId,
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get role")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

func (c *RoleController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateRoleRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("roleId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update role")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoleResponse]{Data: response})
}

func (c *RoleController) Delete(ctx *fiber.Ctx) error {
	roleId := ctx.Params("roleId")

	request := &model.DeleteRoleRequest{
		ID: roleId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete role")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *RoleController) AssignPermissions(ctx *fiber.Ctx) error {
	request := new(model.AssignPermissionsRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.RoleID = ctx.Params("roleId")

	if err := c.UseCase.AssignPermissions(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to assign permissions to role")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

func (c *RoleController) RemovePermissions(ctx *fiber.Ctx) error {
	request := new(model.RemovePermissionsRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.RoleID = ctx.Params("roleId")

	if err := c.UseCase.RemovePermissions(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to remove permissions from role")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
