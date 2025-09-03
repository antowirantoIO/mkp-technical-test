package middleware

import (
	"strings"

	"mkp-boarding-test/internal/model"
	"mkp-boarding-test/internal/service"
	"mkp-boarding-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *usecase.UserUseCase, jwtService service.JWTService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization", "")
		if authorization == "" {
			userUserCase.Log.Warn("Authorization header is missing")
			return fiber.ErrUnauthorized
		}

		// Extract Bearer token
		if !strings.HasPrefix(authorization, "Bearer ") {
			userUserCase.Log.Warn("Authorization header must start with Bearer")
			return fiber.ErrUnauthorized
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		userUserCase.Log.Debugf("JWT Token : %s", token)

		// Validate JWT token
		claims, err := jwtService.ExtractClaimsFromToken(token)
		if err != nil {
			userUserCase.Log.Warnf("Failed to validate JWT token : %+v", err)
			return fiber.ErrUnauthorized
		}

		// Extract user info from claims
		userID, ok := claims["user_id"].(string)
		if !ok {
			userUserCase.Log.Warn("user_id not found in token claims")
			return fiber.ErrUnauthorized
		}

		username, ok := claims["username"].(string)
		if !ok {
			userUserCase.Log.Warn("username not found in token claims")
			return fiber.ErrUnauthorized
		}

		email, ok := claims["email"].(string)
		if !ok {
			userUserCase.Log.Warn("email not found in token claims")
			return fiber.ErrUnauthorized
		}

		auth := &model.Auth{
			ID:       userID,
			Username: username,
			Email:    email,
		}

		userUserCase.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
