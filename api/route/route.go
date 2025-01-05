package route

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/api/middleware"
	"wetees.com/domain"
)

func Setup(
	db *sql.DB,
	app *fiber.App,
	conf *domain.Config,
	logger zerolog.Logger,
	timeout time.Duration,
) {
	var (
		accessTokenSecret = conf.AccessTokenSecret
		secretKey         = conf.SecretKey
		mode              = conf.Mode
	)

	/*
		Public Endpoints
	*/
	publicRoute := app.Group("/api")
	NewWelcomeRoute(conf, publicRoute)
	NewHeatlhCheckRoute(conf, logger, publicRoute)
	NewAuthnRoute(conf, timeout, logger, db, publicRoute)
	NewProductRoute(conf, timeout, logger, db, publicRoute)

	/*
		Protected Endpoints
		User needs to pass bearer token header to authenticate accessing the resources
	*/
	protectedRoute := app.Group("/api")
	protectedRoute.Use(middleware.JwtAuthMiddleware(accessTokenSecret, secretKey, mode, logger))
	NewUserRoute(conf, timeout, logger, db, protectedRoute)
	NewOrderRoute(conf, timeout, logger, db, protectedRoute)
	NewOrderItemRoute(conf, timeout, logger, db, protectedRoute)
	NewBankRoute(conf, timeout, logger, db, protectedRoute)
}
