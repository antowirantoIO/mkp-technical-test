package config

import (
	"golang-clean-architecture/internal/delivery/http"
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/delivery/http/route"
	"golang-clean-architecture/internal/gateway/messaging"
	"golang-clean-architecture/internal/repository"
	"golang-clean-architecture/internal/usecase"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer sarama.SyncProducer
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	contactRepository := repository.NewContactRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)
	roleRepository := repository.NewRoleRepository(config.Log)
	permissionRepository := repository.NewPermissionRepository(config.Log)
	rolePermissionRepository := repository.NewRolePermissionRepository(config.Log)
	operatorRepository := repository.NewOperatorRepository(config.Log)
	shipRepository := repository.NewShipRepository(config.Log)
	harborRepository := repository.NewHarborRepository(config.Log)

	// setup producer
	var userProducer *messaging.UserProducer
	var contactProducer *messaging.ContactProducer
	var addressProducer *messaging.AddressProducer

	if config.Producer != nil {
		userProducer = messaging.NewUserProducer(config.Producer, config.Log)
		contactProducer = messaging.NewContactProducer(config.Producer, config.Log)
		addressProducer = messaging.NewAddressProducer(config.Producer, config.Log)
	}

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer)
	contactUseCase := usecase.NewContactUseCase(config.DB, config.Log, config.Validate, contactRepository, contactProducer)
	addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, config.Validate, contactRepository, addressRepository, addressProducer)
	roleUseCase := usecase.NewRoleUseCase(config.DB, config.Log, config.Validate, roleRepository, permissionRepository, rolePermissionRepository)
	permissionUseCase := usecase.NewPermissionUseCase(config.DB, config.Log, config.Validate, permissionRepository, rolePermissionRepository)
	operatorUseCase := usecase.NewOperatorUseCase(config.DB, config.Log, config.Validate, operatorRepository)
	shipUseCase := usecase.NewShipUseCase(config.DB, config.Log, config.Validate, shipRepository)
	harborUseCase := usecase.NewHarborUseCase(config.DB, config.Log, config.Validate, harborRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	contactController := http.NewContactController(contactUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)
	roleController := http.NewRoleController(roleUseCase, config.Log)
	permissionController := http.NewPermissionController(permissionUseCase, config.Log)
	operatorController := http.NewOperatorController(operatorUseCase, config.Log)
	shipController := http.NewShipController(shipUseCase, config.Log)
	harborController := http.NewHarborController(harborUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                  config.App,
		UserController:       userController,
		ContactController:    contactController,
		AddressController:    addressController,
		RoleController:       roleController,
		PermissionController: permissionController,
		OperatorController:   operatorController,
		ShipController:       shipController,
		HarborController:     harborController,
		AuthMiddleware:       authMiddleware,
	}
	routeConfig.Setup()
	routeConfig.SetupSwaggerRoute()
}
