package http

import (
	"mkp-boarding-test/internal/delivery/http/middleware"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/internal/usecase"
	"mkp-boarding-test/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.RegisterUserRequest true "Register user request"
// @Success 200 {object} model.SwaggerWebResponse "User registered successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /register [post]
func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to register user", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "User registered successfully", response)
}

// Login godoc
// @Summary User login
// @Description Authenticate user with username/email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.LoginUserRequest true "Login user request"
// @Success 200 {object} model.SwaggerWebResponse "Login successful"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Invalid credentials"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /login [post]
func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		return utils.SendErrorResponse(ctx, fiber.StatusUnauthorized, "Invalid credentials", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Login successful", response)
}

// Current godoc
// @Summary Get current user
// @Description Get current authenticated user information
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.SwaggerWebResponse "Current user details"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/users/_current [get]
func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return utils.SendErrorResponse(ctx, fiber.StatusNotFound, "User not found", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Current user retrieved successfully", response)
}

// Logout godoc
// @Summary User logout
// @Description Logout current authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.SwaggerWebResponse "Logout successful"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/users [delete]
func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to logout user", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "Logout successful", response)
}

// Update godoc
// @Summary Update current user
// @Description Update current authenticated user information
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.UpdateUserRequest true "Update user request"
// @Success 200 {object} model.SwaggerWebResponse "User updated successfully"
// @Failure 400 {object} model.SwaggerWebResponse "Bad request"
// @Failure 401 {object} model.SwaggerWebResponse "Unauthorized"
// @Failure 500 {object} model.SwaggerWebResponse "Internal server error"
// @Router /api/users/_current [patch]
func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return utils.SendErrorResponse(ctx, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	request.ID = auth.ID
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		return utils.SendErrorResponse(ctx, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return utils.SendSuccessResponse(ctx, "User updated successfully", response)
}
