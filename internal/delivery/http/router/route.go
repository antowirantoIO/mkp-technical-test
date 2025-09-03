package route

import (
	"mkp-boarding-test/docs"
	"mkp-boarding-test/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type RouteConfig struct {
	App                  *fiber.App
	UserController       *handler.UserController
	RoleController       *handler.RoleController
	PermissionController *handler.PermissionController
	OperatorController   *handler.OperatorController
	ShipController       *handler.ShipController
	HarborController     *handler.HarborController
	AuthMiddleware       fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupSwaggerRoute()
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupSwaggerRoute() {
	// Initialize swagger docs
	_ = docs.SwaggerInfo

	// Swagger documentation route
	c.App.Get("/swagger/*", swagger.HandlerDefault)
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/register", c.UserController.Register)
	c.App.Post("/login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	// Create API group with auth middleware
	api := c.App.Group("/api", c.AuthMiddleware)

	// User routes
	api.Delete("/users", c.UserController.Logout)
	api.Patch("/users/_current", c.UserController.Update)
	api.Get("/users/_current", c.UserController.Current)

	// Role routes
	api.Get("/roles", c.RoleController.List)
	api.Post("/roles", c.RoleController.Create)
	api.Put("/roles/:roleId", c.RoleController.Update)
	api.Get("/roles/:roleId", c.RoleController.Get)
	api.Delete("/roles/:roleId", c.RoleController.Delete)
	api.Post("/roles/:roleId/permissions", c.RoleController.AssignPermissions)
	api.Delete("/roles/:roleId/permissions", c.RoleController.RemovePermissions)

	// Permission routes
	api.Get("/permissions", c.PermissionController.List)
	api.Post("/permissions", c.PermissionController.Create)
	api.Put("/permissions/:permissionId", c.PermissionController.Update)
	api.Get("/permissions/:permissionId", c.PermissionController.Get)
	api.Delete("/permissions/:permissionId", c.PermissionController.Delete)

	// Operator routes
	api.Get("/operators", c.OperatorController.List)
	api.Post("/operators", c.OperatorController.Create)
	api.Put("/operators/:operatorId", c.OperatorController.Update)
	api.Get("/operators/:operatorId", c.OperatorController.Get)
	api.Delete("/operators/:operatorId", c.OperatorController.Delete)

	// Ship routes
	api.Get("/ships", c.ShipController.List)
	api.Post("/ships", c.ShipController.Create)
	api.Put("/ships/:shipId", c.ShipController.Update)
	api.Get("/ships/:shipId", c.ShipController.Get)
	api.Delete("/ships/:shipId", c.ShipController.Delete)

	// Harbor routes
	api.Get("/harbors", c.HarborController.List)
	api.Post("/harbors", c.HarborController.Create)
	api.Put("/harbors/:harborId", c.HarborController.Update)
	api.Get("/harbors/:harborId", c.HarborController.Get)
	api.Delete("/harbors/:harborId", c.HarborController.Delete)
}
