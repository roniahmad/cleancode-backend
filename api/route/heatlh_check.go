package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/response"
)

// Heatlh check handler
func NewHeatlhCheckRoute(conf *domain.Config, logger zerolog.Logger, app fiber.Router) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return response.SendSuccessResponse(c, conf.Mode, logger, "OK")
	})
}
