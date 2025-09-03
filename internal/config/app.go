package config

import (
	"time"

	"mkp-boarding-test/internal/delivery/http"
	"mkp-boarding-test/internal/delivery/http/middleware"
	"mkp-boarding-test/internal/delivery/http/route"
	"mkp-boarding-test/internal/gateway/messaging"
	"mkp-boarding-test/internal/repository"
	"mkp-boarding-test/internal/service"
	"mkp-boarding-test/internal/usecase"

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
	roleRepository := repository.NewRoleRepository(config.Log)
	permissionRepository := repository.NewPermissionRepository(config.Log)
	rolePermissionRepository := repository.NewRolePermissionRepository(config.Log)
	operatorRepository := repository.NewOperatorRepository(config.Log)
	shipRepository := repository.NewShipRepository(config.Log)
	harborRepository := repository.NewHarborRepository(config.Log)

	// setup JWT service
	jwtService := service.NewJWTService(
		"your-secret-key",
		"your-refresh-key",
		24*time.Hour,  // token expiry
		7*24*time.Hour, // refresh token expiry
	)

	// setup producer
	var userProducer *messaging.UserProducer

	if config.Producer != nil {
		userProducer = messaging.NewUserProducer(config.Producer, config.Log)
	}

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer, jwtService)
	roleUseCase := usecase.NewRoleUseCase(config.DB, config.Log, config.Validate, roleRepository, permissionRepository, rolePermissionRepository)
	permissionUseCase := usecase.NewPermissionUseCase(config.DB, config.Log, config.Validate, permissionRepository, rolePermissionRepository)
	operatorUseCase := usecase.NewOperatorUseCase(config.DB, config.Log, config.Validate, operatorRepository)
	shipUseCase := usecase.NewShipUseCase(config.DB, config.Log, config.Validate, shipRepository)
	harborUseCase := usecase.NewHarborUseCase(config.DB, config.Log, config.Validate, harborRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	roleController := http.NewRoleController(roleUseCase, config.Log)
	permissionController := http.NewPermissionController(permissionUseCase, config.Log)
	operatorController := http.NewOperatorController(operatorUseCase, config.Log)
	shipController := http.NewShipController(shipUseCase, config.Log)
	harborController := http.NewHarborController(harborUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase, jwtService)

	routeConfig := route.RouteConfig{
		App:                  config.App,
		UserController:       userController,
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
