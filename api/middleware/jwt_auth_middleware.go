package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/internal/jwt"
	"wetees.com/internal/response"
	"wetees.com/internal/vars"
)

func JwtAuthMiddleware(key, secret, mode string, log zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		xAuthHeader := c.Get(fiber.HeaderAuthorization)
		token, err := jwt.GetAuthToken(xAuthHeader)
		if err != nil {
			return response.SendErrorResponse(c, fiber.StatusUnauthorized, mode, log, vars.ErrUnAuthorized)
		}

		authorized, err := jwt.IsAuthorized(token, key)
		if authorized {
			claims, err := jwt.ExtractClaimsFromToken(token, key)
			if err != nil {
				return response.SendErrorResponse(c, fiber.StatusUnauthorized, mode, log, err)
			}
			c.Set("x-user-id", claims["id"].(string))
			return c.Next()
		}

		return response.SendErrorResponse(c, fiber.StatusUnauthorized, mode, log, vars.ErrUnAuthorized)
	}
}
