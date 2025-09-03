package config

import (
	"time"

	"mkp-boarding-test/internal/delivery/http/handler"
	"mkp-boarding-test/internal/delivery/http/middleware"
	route "mkp-boarding-test/internal/delivery/http/router"
	"mkp-boarding-test/internal/gateway/messaging"
	harborRepo "mkp-boarding-test/internal/infrastructure/repository/harbor"
	operatorRepo "mkp-boarding-test/internal/infrastructure/repository/operator"
	permissionRepo "mkp-boarding-test/internal/infrastructure/repository/permission"
	roleRepo "mkp-boarding-test/internal/infrastructure/repository/role"

	harborUsecase "mkp-boarding-test/internal/application/usecase/harbor"
	operatorUsecase "mkp-boarding-test/internal/application/usecase/operator"
	permissionUsecase "mkp-boarding-test/internal/application/usecase/permission"
	roleUsecase "mkp-boarding-test/internal/application/usecase/role"
	shipUsecase "mkp-boarding-test/internal/application/usecase/ship"
	userUsecase "mkp-boarding-test/internal/application/usecase/user"
	shipRepo "mkp-boarding-test/internal/infrastructure/repository/ship"
	userRepo "mkp-boarding-test/internal/infrastructure/repository/user"
	"mkp-boarding-test/pkg/service"

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
	userRepository := userRepo.NewUserRepository(config.Log)
	roleRepository := roleRepo.NewRoleRepository(config.Log)
	permissionRepository := permissionRepo.NewPermissionRepository(config.Log)

	operatorRepository := operatorRepo.NewOperatorRepository(config.Log)
	shipRepository := shipRepo.NewShipRepository(config.Log)
	harborRepository := harborRepo.NewHarborRepository(config.Log)

	// setup JWT service
	jwtService := service.NewJWTService(
		"your-secret-key",
		"your-refresh-key",
		24*time.Hour,   // token expiry
		7*24*time.Hour, // refresh token expiry
	)

	// setup producer
	var userProducer *messaging.UserProducer

	if config.Producer != nil {
		userProducer = messaging.NewUserProducer(config.Producer, config.Log)
	}

	// setup use cases
	userUseCase := userUsecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer, jwtService)
	roleUseCase := roleUsecase.NewRoleUseCase(config.DB, config.Log, config.Validate, roleRepository, permissionRepository)
	permissionUseCase := permissionUsecase.NewPermissionUseCase(config.DB, config.Log, config.Validate, permissionRepository)
	operatorUseCase := operatorUsecase.NewOperatorUseCase(config.DB, config.Log, config.Validate, operatorRepository)
	shipUseCase := shipUsecase.NewShipUseCase(config.DB, config.Log, config.Validate, shipRepository)
	harborUseCase := harborUsecase.NewHarborUseCase(config.DB, config.Log, config.Validate, harborRepository)

	// setup controller
	userController := handler.NewUserController(userUseCase, config.Log)
	roleController := handler.NewRoleController(roleUseCase, config.Log)
	permissionController := handler.NewPermissionController(permissionUseCase, config.Log)
	operatorController := handler.NewOperatorController(operatorUseCase, config.Log)
	shipController := handler.NewShipController(shipUseCase, config.Log)
	harborController := handler.NewHarborController(harborUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase, jwtService, config.Log)

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
