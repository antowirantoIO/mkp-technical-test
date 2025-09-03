package route

import (
	"golang-clean-architecture/docs"
	"golang-clean-architecture/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type RouteConfig struct {
	App                  *fiber.App
	UserController       *http.UserController
	ContactController    *http.ContactController
	AddressController    *http.AddressController
	RoleController       *http.RoleController
	PermissionController *http.PermissionController
	OperatorController   *http.OperatorController
	ShipController       *http.ShipController
	HarborController     *http.HarborController
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
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	
	// User routes
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)

	// Contact routes
	c.App.Get("/api/contacts", c.ContactController.List)
	c.App.Post("/api/contacts", c.ContactController.Create)
	c.App.Put("/api/contacts/:contactId", c.ContactController.Update)
	c.App.Get("/api/contacts/:contactId", c.ContactController.Get)
	c.App.Delete("/api/contacts/:contactId", c.ContactController.Delete)

	// Address routes
	c.App.Get("/api/contacts/:contactId/addresses", c.AddressController.List)
	c.App.Post("/api/contacts/:contactId/addresses", c.AddressController.Create)
	c.App.Put("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	c.App.Get("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	c.App.Delete("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)

	// Role routes
	c.App.Get("/api/roles", c.RoleController.List)
	c.App.Post("/api/roles", c.RoleController.Create)
	c.App.Put("/api/roles/:roleId", c.RoleController.Update)
	c.App.Get("/api/roles/:roleId", c.RoleController.Get)
	c.App.Delete("/api/roles/:roleId", c.RoleController.Delete)
	c.App.Post("/api/roles/:roleId/permissions", c.RoleController.AssignPermissions)
	c.App.Delete("/api/roles/:roleId/permissions", c.RoleController.RemovePermissions)

	// Permission routes
	c.App.Get("/api/permissions", c.PermissionController.List)
	c.App.Post("/api/permissions", c.PermissionController.Create)
	c.App.Put("/api/permissions/:permissionId", c.PermissionController.Update)
	c.App.Get("/api/permissions/:permissionId", c.PermissionController.Get)
	c.App.Delete("/api/permissions/:permissionId", c.PermissionController.Delete)

	// Operator routes
	c.App.Get("/api/operators", c.OperatorController.List)
	c.App.Post("/api/operators", c.OperatorController.Create)
	c.App.Put("/api/operators/:operatorId", c.OperatorController.Update)
	c.App.Get("/api/operators/:operatorId", c.OperatorController.Get)
	c.App.Delete("/api/operators/:operatorId", c.OperatorController.Delete)

	// Ship routes
	c.App.Get("/api/ships", c.ShipController.List)
	c.App.Post("/api/ships", c.ShipController.Create)
	c.App.Put("/api/ships/:shipId", c.ShipController.Update)
	c.App.Get("/api/ships/:shipId", c.ShipController.Get)
	c.App.Delete("/api/ships/:shipId", c.ShipController.Delete)

	// Harbor routes
	c.App.Get("/api/harbors", c.HarborController.List)
	c.App.Post("/api/harbors", c.HarborController.Create)
	c.App.Put("/api/harbors/:harborId", c.HarborController.Update)
	c.App.Get("/api/harbors/:harborId", c.HarborController.Get)
	c.App.Delete("/api/harbors/:harborId", c.HarborController.Delete)
}
