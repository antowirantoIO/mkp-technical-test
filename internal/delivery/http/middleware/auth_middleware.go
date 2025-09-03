package middleware

import (
	"strings"

	"mkp-boarding-test/internal/domain/usecase"
	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/pkg/service"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase usecase.UserUseCase, jwtService service.JWTService, log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization", "")
		if authorization == "" {
			log.Warn("Authorization header is missing")
			return fiber.ErrUnauthorized
		}

		// Extract Bearer token
		if !strings.HasPrefix(authorization, "Bearer ") {
			log.Warn("Authorization header must start with Bearer")
			return fiber.ErrUnauthorized
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		log.Debugf("JWT Token : %s", token)

		// Validate JWT token
		claims, err := jwtService.ExtractClaimsFromToken(token)
		if err != nil {
			log.Warnf("Failed to validate JWT token : %+v", err)
			return fiber.ErrUnauthorized
		}

		// Extract user info from claims
		userID, ok := claims["user_id"].(string)
		if !ok {
			log.Warn("user_id not found in token claims")
			return fiber.ErrUnauthorized
		}

		username, ok := claims["username"].(string)
		if !ok {
			log.Warn("username not found in token claims")
			return fiber.ErrUnauthorized
		}

		email, ok := claims["email"].(string)
		if !ok {
			log.Warn("email not found in token claims")
			return fiber.ErrUnauthorized
		}

		auth := &model.Auth{
			ID:       userID,
			Username: username,
			Email:    email,
		}

		log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
