package middleware

import (
	"go-api/utils/jwtutil"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		claims, err := jwtutil.VerifyToken(accessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON("Unauthorized")
		}

		c.Locals("users", claims)

		return c.Next()
	}
}
